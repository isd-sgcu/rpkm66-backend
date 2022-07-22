package group

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	baanModel "github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	baan_group_selection "github.com/isd-sgcu/rnkm65-backend/src/app/model/baan-group-selection"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/service/baan"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	repo               IRepository
	baanRepo           IBaanRepository
	userRepo           IUserRepository
	cacheRepo          ICacheRepository
	fileSrv            IFileService
	conf               config.App
	baanGroupSelection IBaanGroupSelectRepository
}

type IRepository interface {
	FindGroupByToken(string, *group.Group) error
	FindGroupById(string, *group.Group) error
	FindGroupWithBaans(string, *group.Group) error
	Create(*group.Group) error
	UpdateWithLeader(string, *group.Group) error
	RemoveAllBaan(g *group.Group) error
	Delete(string) error
}

type ICacheRepository interface {
	SaveCache(key string, value interface{}, ttl int) error
	GetCache(key string, value interface{}) error
}

type IBaanGroupSelectRepository interface {
	SaveBaansSelection(*[]*baan_group_selection.BaanGroupSelection) error
	FindBaans(string, *[]*baan_group_selection.BaanGroupSelection) error
}

type IBaanRepository interface {
	FindMulti([]string, *[]*baanModel.Baan) error
}

type IUserRepository interface {
	FindOne(string, *user.User) error
	Update(string, *user.User) error
}

func NewService(repo IRepository, userRepo IUserRepository, baanGroupSelectionRepo IBaanGroupSelectRepository, fileSrv IFileService, cacheRepo ICacheRepository, baanRepo IBaanRepository, conf config.App) *Service {
	return &Service{
		repo:               repo,
		userRepo:           userRepo,
		cacheRepo:          cacheRepo,
		fileSrv:            fileSrv,
		baanRepo:           baanRepo,
		baanGroupSelection: baanGroupSelectionRepo,
		conf:               conf,
	}
}

type IFileService interface {
	GetSignedUrl(string) (string, error)
}

func (s *Service) FindOne(_ context.Context, req *proto.FindOneGroupRequest) (res *proto.FindOneGroupResponse, err error) {
	_, err = uuid.Parse(req.UserId)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find one").
			Str("user_id", req.UserId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	usr := &user.User{}
	err = s.userRepo.FindOne(req.UserId, usr)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find one").
			Str("user_id", req.UserId).
			Msg("Not found user")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	//if the user is not in the group -> create group
	if usr.GroupID == nil {
		newGrp := &group.Group{
			LeaderID: usr.ID.String(),
		}
		err = s.repo.Create(newGrp)
		if err != nil {
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "find one").
				Str("student_id", usr.StudentID).
				Msg("Fail to create group")
			return nil, status.Error(codes.Internal, "failed to create group")
		}

		usr.GroupID = utils.UUIDAdr(newGrp.ID)
		_ = s.userRepo.Update(usr.ID.String(), usr)

		updateGrp := &group.Group{}
		_ = s.repo.FindGroupByToken(newGrp.Token, updateGrp)

		grpDto, err := s.RawToDto(updateGrp)
		if err != nil {
			return nil, err
		}

		log.Info().
			Str("service", "group").
			Str("module", "find one").
			Str("student_id", usr.StudentID).
			Msg("Find group success")

		return &proto.FindOneGroupResponse{Group: grpDto}, nil
	}

	grp := &group.Group{}
	err = s.repo.FindGroupById((*usr.GroupID).String(), grp)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find one").
			Str("student_id", usr.StudentID).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}

	var baanInfos []*proto.BaanInfo
	err = s.cacheRepo.GetCache(usr.GroupID.String(), &baanInfos)
	if err != nil {
		if err != redis.Nil {
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "find one").
				Str("student_id", usr.StudentID).
				Msg("Cannot connect to redis")
			return nil, status.Error(codes.Internal, "Cannot connect to redis")
		}

		grp := &group.Group{}
		err := s.repo.FindGroupWithBaans(usr.GroupID.String(), grp)
		if err != nil {
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "select baan").
				Str("user_id", req.UserId).
				Msg("Not found group")
			return nil, status.Error(codes.NotFound, "Not found group")
		}

		var baans []*baan_group_selection.BaanGroupSelection
		err = s.baanGroupSelection.FindBaans(grp.ID.String(), &baans)
		if err != nil {
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "select baan").
				Str("user_id", req.UserId).
				Msg("Not found baan")
			return nil, status.Error(codes.NotFound, "Not found baan")
		}

		baanInfos = baan.RawToDtoBaanSelection(&baans)
		err = s.cacheRepo.SaveCache(grp.ID.String(), &baanInfos, s.conf.BaanCacheTTL)
		if err != nil {
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "find one").
				Str("student_id", usr.StudentID).
				Msg("Cannot connect to redis")
			return nil, status.Error(codes.Internal, "Cannot connect to redis")
		}
	}

	grpDto, err := s.RawToDto(grp)
	if err != nil {
		return nil, err
	}

	grpDto.Baans = baanInfos

	log.Info().
		Str("service", "group").
		Str("module", "find one").
		Str("student_id", usr.StudentID).
		Msg("Find group success")

	return &proto.FindOneGroupResponse{Group: grpDto}, nil
}

func (s *Service) FindByToken(_ context.Context, req *proto.FindByTokenGroupRequest) (res *proto.FindByTokenGroupResponse, err error) {
	raw := &group.Group{}

	err = s.repo.FindGroupByToken(req.Token, raw)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find group by token").
			Str("token", req.Token).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}

	usr := &user.User{}

	err = s.userRepo.FindOne(raw.LeaderID, usr)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find group by token").
			Str("token", req.Token).
			Msg("Not found user")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	log.Info().
		Str("service", "group").
		Str("module", "find group by token").
		Str("token", req.Token).
		Msg("Find group by token success")

	leaderUrl, err := s.GetUserImage(usr)
	if err != nil {
		return nil, err
	}
	return &proto.FindByTokenGroupResponse{
		Id:    raw.ID.String(),
		Token: raw.Token,
		Leader: &proto.UserInfo{
			Id:        usr.ID.String(),
			Firstname: usr.Firstname,
			Lastname:  usr.Lastname,
			ImageUrl:  leaderUrl,
		},
	}, nil
}

func (s *Service) Update(_ context.Context, req *proto.UpdateGroupRequest) (res *proto.UpdateGroupResponse, err error) {
	raw, err := DtoToRaw(req.Group)

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "update").
			Str("user_id", req.LeaderId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid group id")
	}

	err = s.repo.UpdateWithLeader(req.LeaderId, raw)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "update").
			Str("user_id", req.LeaderId).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}

	grpDto, err := s.RawToDto(raw)
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("service", "group").
		Str("module", "update").
		Str("user_id", req.LeaderId).
		Msg("Update group success")
	return &proto.UpdateGroupResponse{Group: grpDto}, nil
}

func (s *Service) Join(_ context.Context, req *proto.JoinGroupRequest) (res *proto.JoinGroupResponse, err error) {
	//check whether the user id is valid or not
	if _, err = uuid.Parse(req.UserId); err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	//Get group to check whether the joined group is existed or not
	joinGroup := &group.Group{}
	err = s.repo.FindGroupByToken(req.Token, joinGroup)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}
	//check if the group is fulled or not
	if len(joinGroup.Members) >= 3 {
		return nil, status.Error(codes.PermissionDenied, "group full")
	}

	//Find user if user is existed
	joinUser := &user.User{}
	err = s.userRepo.FindOne(req.UserId, joinUser)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Not found user")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	//check if the joining user is in the joined group or not
	if (*joinUser.GroupID).String() == joinGroup.ID.String() {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Not allowed")
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	//Get group of joining user to check whether the user is leader or not
	prevGrp := &group.Group{}
	err = s.repo.FindGroupById((*joinUser.GroupID).String(), prevGrp)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}

	if req.UserId == prevGrp.LeaderID && len(prevGrp.Members) > 1 {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Not allowed")
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	prevGroupId := prevGrp.ID.String()
	joinUser.GroupID = &joinGroup.ID
	//update user
	err = s.userRepo.Update(joinUser.ID.String(), joinUser)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Fail to Update user")
		return nil, status.Error(codes.NotFound, "fail to update user")
	}

	//if the joining user is the leader and there is only one in the group -> delete the previous group
	if req.UserId == prevGrp.LeaderID && len(prevGrp.Members) == 1 {
		_ = s.repo.Delete(prevGroupId)
	}

	newGrp := &group.Group{}
	_ = s.repo.FindGroupByToken(req.Token, newGrp)

	grpDto, err := s.RawToDto(newGrp)
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("service", "group").
		Str("module", "join").
		Str("student_id", joinUser.StudentID).
		Msg("Join group Success")
	return &proto.JoinGroupResponse{Group: grpDto}, nil
}

func (s *Service) DeleteMember(_ context.Context, req *proto.DeleteMemberGroupRequest) (res *proto.DeleteMemberGroupResponse, err error) {
	_, err = uuid.Parse(req.UserId)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "delete member").
			Str("user_id", req.LeaderId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	removedUser := &user.User{}
	err = s.userRepo.FindOne(req.UserId, removedUser)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "delete member").
			Str("user_id", req.LeaderId).
			Msg("Not found user")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	deletedGrp := &group.Group{}
	err = s.repo.FindGroupById((*removedUser.GroupID).String(), deletedGrp)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "delete member").
			Str("user_id", req.LeaderId).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}

	if deletedGrp.LeaderID != req.LeaderId {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "delete member").
			Str("user_id", req.LeaderId).
			Msg("Not allowed")
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	//create a new group for removed user
	newGroup := &group.Group{
		LeaderID: req.UserId,
	}
	err = s.repo.Create(newGroup)
	removedUser.GroupID = &newGroup.ID

	_ = s.userRepo.Update(removedUser.ID.String(), removedUser)

	afterDeleteGrp := &group.Group{}
	_ = s.repo.FindGroupByToken(deletedGrp.Token, afterDeleteGrp)

	grpDto, err := s.RawToDto(afterDeleteGrp)
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("service", "group").
		Str("module", "delete member").
		Str("user_id", req.LeaderId).
		Msg("Delete member success")
	return &proto.DeleteMemberGroupResponse{Group: grpDto}, nil
}

func (s *Service) Leave(_ context.Context, req *proto.LeaveGroupRequest) (res *proto.LeaveGroupResponse, err error) {
	if _, err = uuid.Parse(req.UserId); err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "leave").
			Str("user_id", req.UserId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	leavedUser := &user.User{}
	err = s.userRepo.FindOne(req.UserId, leavedUser)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "leave").
			Str("user_id", req.UserId).
			Msg("Not found user")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	prevGrp := &group.Group{}
	err = s.repo.FindGroupById((*leavedUser.GroupID).String(), prevGrp)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "leave").
			Str("student_id", leavedUser.StudentID).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}

	if req.UserId == prevGrp.LeaderID {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "leave").
			Str("student_id", leavedUser.StudentID).
			Msg("Not allowed")
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	in := &group.Group{
		LeaderID: leavedUser.ID.String(),
	}
	err = s.repo.Create(in)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "leave").
			Str("student_id", leavedUser.StudentID).
			Msg("Fail to create group")
		return nil, status.Error(codes.Internal, "failed to create group")
	}
	leavedUser.GroupID = &in.ID

	err = s.userRepo.Update(leavedUser.ID.String(), leavedUser)

	newGroup := &group.Group{}

	_ = s.repo.FindGroupByToken(in.Token, newGroup)

	grpDto, err := s.RawToDto(newGroup)
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("service", "group").
		Str("module", "leave").
		Str("student_id", leavedUser.StudentID).
		Msg("Leave group success")
	return &proto.LeaveGroupResponse{Group: grpDto}, nil
}

func (s *Service) SelectBaan(_ context.Context, req *proto.SelectBaanRequest) (res *proto.SelectBaanResponse, err error) {
	if len(req.Baans) != s.conf.NBaan {
		log.Error().
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Invalid number of baan")
		return nil, status.Error(codes.InvalidArgument, "invalid numbers of baan")
	}

	if utils.IsDuplicatedString(req.Baans) {
		log.Error().
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Duplicated baan")
		return nil, status.Error(codes.InvalidArgument, "duplicated baan")
	}

	leader := &user.User{}
	err = s.userRepo.FindOne(req.UserId, leader)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found user")
	}

	grp := &group.Group{}
	err = s.repo.FindGroupWithBaans(leader.GroupID.String(), grp)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found group")
	}

	if leader.ID.String() != grp.LeaderID {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Forbidden action not leader select baan")
		return nil, status.Error(codes.PermissionDenied, "Insufficiency permission")
	}

	var baanSelections []*baan_group_selection.BaanGroupSelection
	for i, bId := range req.Baans {
		bId, err := uuid.Parse(bId)
		if err != nil {
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "select baan").
				Str("user_id", req.UserId).
				Msg("Duplicated baan")
			return nil, status.Error(codes.Internal, "Cannot parse group id to int")
		}

		baanSelect := baan_group_selection.BaanGroupSelection{
			BaanID:  &bId,
			GroupID: &grp.ID,
			Order:   i + 1,
		}

		baanSelections = append(baanSelections, &baanSelect)
	}

	err = s.repo.RemoveAllBaan(&group.Group{Base: model.Base{ID: grp.ID}})
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Error while clearing the baan selection")
		return nil, status.Error(codes.Internal, "Error while clearing the baan selection")
	}

	err = s.baanGroupSelection.SaveBaansSelection(&baanSelections)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Not found baan")
		return nil, status.Error(codes.NotFound, "Not found baan")
	}

	var baans []*baan_group_selection.BaanGroupSelection
	err = s.baanGroupSelection.FindBaans(grp.ID.String(), &baans)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Not found baan")
		return nil, status.Error(codes.NotFound, "Not found baan")
	}

	baanCache := baan.RawToDtoBaanSelection(&baans)
	err = s.cacheRepo.SaveCache(grp.ID.String(), &baanCache, s.conf.BaanCacheTTL)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "select baan").
			Str("user_id", req.UserId).
			Msg("Cannot connect to redis")
		return nil, status.Error(codes.Internal, "Cannot connect to redis")
	}

	log.Info().
		Str("service", "group").
		Str("module", "select baan").
		Str("user_id", req.UserId).
		Msg("Successfully update baan selection")

	return &proto.SelectBaanResponse{Success: true}, nil
}

func DtoToRaw(in *proto.Group) (result *group.Group, err error) {
	var members []*user.User
	for _, usr := range in.Members {
		id, err := uuid.Parse(usr.Id)
		if err != nil {
			return nil, err
		}

		newUser := &user.User{
			Base: model.Base{
				ID:        id,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Firstname: usr.Firstname,
			Lastname:  usr.Lastname,
		}
		members = append(members, newUser)
	}

	var id uuid.UUID
	if in.Id != "" {
		id, err = uuid.Parse(in.Id)
		if err != nil {
			return nil, err
		}
	}
	return &group.Group{
		Base: model.Base{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: in.LeaderID,
		Token:    in.Token,
		Members:  members,
	}, nil
}

func (s *Service) RawToDto(in *group.Group) (result *proto.Group, err error) {
	members, err := s.GetMembersImages(in.Members)
	if err != nil {
		return nil, err
	}
	grp := &proto.Group{
		Id:       in.ID.String(),
		LeaderID: in.LeaderID,
		Token:    in.Token,
		Members:  members,
	}
	return grp, nil
}

func (s *Service) GetMembersImages(users []*user.User) (result []*proto.UserInfo, err error) {
	var members []*proto.UserInfo
	for _, usr := range users {
		url, err := s.GetUserImage(usr)
		if err != nil {
			return nil, err
		}
		newUser := &proto.UserInfo{
			Id:        usr.ID.String(),
			Firstname: usr.Firstname,
			Lastname:  usr.Lastname,
			ImageUrl:  url,
		}
		members = append(members, newUser)
	}

	return members, nil
}

func (s *Service) GetUserImage(usr *user.User) (result string, err error) {
	url, err := s.fileSrv.GetSignedUrl(usr.ID.String())
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				log.Error().
					Err(err).
					Str("service", "group").
					Str("module", "find one").
					Str("student_id", usr.StudentID).
					Msg("Not found image")
				return "", nil
			case codes.Unavailable:
				log.Error().
					Err(err).
					Str("service", "group").
					Str("module", "find one").
					Str("student_id", usr.StudentID).
					Msg("Fail to connect database")
				return "", status.Error(codes.Unavailable, "fail to connect database")

			default:
				log.Error().
					Err(err).
					Str("service", "group").
					Str("module", "find one").
					Str("student_id", usr.StudentID).
					Msg("Error while connecting to service")
				return "", err
			}
		}
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find one").
			Str("student_id", usr.StudentID).
			Msg("Error while connecting to service")
		return "", err
	}
	return url, nil
}
