package group

import (
	"context"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	repo     IRepository
	userRepo IUserRepository
	fileSrv  IFileService
}

type IRepository interface {
	FindGroupByToken(string, *group.Group) error
	FindGroupById(string, *group.Group) error
	Create(*group.Group) error
	UpdateWithLeader(string, *group.Group) error
	Delete(string) error
}

type IUserRepository interface {
	FindOne(string, *user.User) error
	Update(string, *user.User) error
}

type IFileService interface {
	GetSignedUrl(string) (string, error)
}

func NewService(repo IRepository, userRepo IUserRepository, fileSrv IFileService) *Service {
	return &Service{repo: repo, userRepo: userRepo, fileSrv: fileSrv}
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

		grpDto, err := RawToDto(s, updateGrp)
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

	grpDto, err := RawToDto(s, grp)
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
	return &proto.FindByTokenGroupResponse{
		Id:              raw.ID.String(),
		LeaderID:        raw.LeaderID,
		Token:           raw.Token,
		LeaderFirstName: usr.Firstname,
		LeaderLastName:  usr.Lastname,
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

	grpDto, err := RawToDto(s, raw)
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

	grpDto, err := RawToDto(s, newGrp)
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

	grpDto, err := RawToDto(s, afterDeleteGrp)
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

	grpDto, err := RawToDto(s, newGroup)
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
			Firstname: usr.FirstName,
			Lastname:  usr.LastName,
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

func RawToDto(s *Service, in *group.Group) (result *proto.Group, err error) {
	var members []*proto.UserInfo
	foundAll := true
	for _, usr := range in.Members {
		url, err := s.fileSrv.GetSignedUrl(usr.ID.String())
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				switch st.Code() {
				case codes.NotFound:
					foundAll = false
				case codes.Unavailable:
					log.Error().
						Err(err).
						Str("service", "group").
						Str("module", "find one").
						Msg("Fail to connect database")
					return nil, status.Error(codes.Unavailable, "fail to connect database")

				default:
					log.Error().
						Err(err).
						Str("service", "group").
						Str("module", "find one").
						Msg("Error while connecting to service")
					return nil, err
				}
			}
			log.Error().
				Err(err).
				Str("service", "group").
				Str("module", "find one").
				Msg("Error while connecting to service")
			if foundAll {
				return nil, err
			}
		}

		newUser := &proto.UserInfo{
			Id:        usr.ID.String(),
			FirstName: usr.Firstname,
			LastName:  usr.Lastname,
			ImageUrl:  url,
		}
		members = append(members, newUser)
	}

	grp := &proto.Group{
		Id:       in.ID.String(),
		LeaderID: in.LeaderID,
		Token:    in.Token,
		Members:  members,
	}
	if !foundAll {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find one").
			Msg("Not found image url")
	}
	return grp, nil
}
