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

func NewService(repo IRepository, userRepo IUserRepository) *Service {
	return &Service{repo: repo, userRepo: userRepo}
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

	log.Info().
		Str("service", "group").
		Str("module", "find one").
		Str("student_id", usr.StudentID).
		Msg("Find group success")
	return &proto.FindOneGroupResponse{Group: RawToDto(grp)}, nil
}

func (s *Service) FindByToken(_ context.Context, req *proto.FindByTokenGroupRequest) (res *proto.FindByTokenGroupResponse, err error) {
	raw := group.Group{}

	err = s.repo.FindGroupByToken(req.Token, &raw)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "find group by token").
			Str("token", req.Token).
			Msg("Not found group")
		return nil, status.Error(codes.NotFound, "group not found")
	}
	log.Info().
		Str("service", "group").
		Str("module", "find group by token").
		Str("token", req.Token).
		Msg("Find group by token success")
	return &proto.FindByTokenGroupResponse{Group: RawToDto(&raw)}, nil
}

func (s *Service) Create(_ context.Context, req *proto.CreateGroupRequest) (res *proto.CreateGroupResponse, err error) {
	if _, err = uuid.Parse(req.UserId); err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "create").
			Str("user_id", req.UserId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid leader id")
	}

	usr := &user.User{}
	err = s.userRepo.FindOne(req.UserId, usr)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "create").
			Str("user_id", req.UserId).
			Msg("Not found user")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	in := &group.Group{
		LeaderID: req.UserId,
	}
	err = s.repo.Create(in)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "create").
			Str("student_id", usr.StudentID).
			Msg("Fail to create group")
		return nil, status.Error(codes.Internal, "failed to create group")
	}

	usr.GroupID = &in.ID
	err = s.userRepo.Update(usr.ID.String(), usr)
	in.Members = []*user.User{usr}

	log.Info().
		Str("service", "group").
		Str("module", "create").
		Str("student_id", usr.StudentID).
		Msg("Create group success")
	return &proto.CreateGroupResponse{Group: RawToDto(in)}, nil
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

	log.Info().
		Str("service", "group").
		Str("module", "update").
		Str("user_id", req.LeaderId).
		Msg("Update group success")
	return &proto.UpdateGroupResponse{Group: RawToDto(raw)}, nil
}

func (s *Service) Join(_ context.Context, req *proto.JoinGroupRequest) (res *proto.JoinGroupResponse, err error) {
	if req.IsLeader && req.Members > 1 {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Not allowed")
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	if _, err = uuid.Parse(req.UserId); err != nil {
		log.Error().
			Err(err).
			Str("service", "group").
			Str("module", "join").
			Str("user_id", req.UserId).
			Msg("Invalid user id")
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

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

	if len(joinGroup.Members) >= 3 {
		return nil, status.Error(codes.PermissionDenied, "group full")
	}

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
	prevGroupId := joinUser.GroupID
	joinUser.GroupID = &joinGroup.ID
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

	if req.IsLeader && req.Members == 1 {
		_ = s.repo.Delete(prevGroupId.String())
	}

	joinGroup.Members = append(joinGroup.Members, joinUser)

	log.Info().
		Str("service", "group").
		Str("module", "join").
		Str("student_id", joinUser.StudentID).
		Msg("Join group Success")
	return &proto.JoinGroupResponse{Group: RawToDto(joinGroup)}, nil
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
	err = s.repo.FindGroupById(removedUser.GroupID.String(), deletedGrp)
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

	log.Info().
		Str("service", "group").
		Str("module", "delete member").
		Str("user_id", req.LeaderId).
		Msg("Delete member success")
	return &proto.DeleteMemberGroupResponse{Group: RawToDto(afterDeleteGrp)}, nil
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
	err = s.repo.FindGroupById(leavedUser.GroupID.String(), prevGrp)
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

	log.Info().
		Str("service", "group").
		Str("module", "leave").
		Str("student_id", leavedUser.StudentID).
		Msg("Leave group success")
	return &proto.LeaveGroupResponse{Group: RawToDto(newGroup)}, nil
}

func DtoToRaw(in *proto.Group) (result *group.Group, err error) {
	var members []*user.User
	for _, usr := range in.Members {
		id, err := uuid.Parse(usr.Id)
		if err != nil {
			return nil, err
		}
		groupId, err := uuid.Parse(usr.GroupId)
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
			Title:           usr.Title,
			Firstname:       usr.Firstname,
			Lastname:        usr.Lastname,
			Nickname:        usr.Nickname,
			StudentID:       usr.StudentID,
			Faculty:         usr.Faculty,
			Year:            usr.Year,
			Phone:           usr.Phone,
			LineID:          usr.LineID,
			Email:           usr.Email,
			AllergyFood:     usr.AllergyFood,
			FoodRestriction: usr.FoodRestriction,
			AllergyMedicine: usr.AllergyMedicine,
			Disease:         usr.Disease,
			CanSelectBaan:   &usr.CanSelectBaan,
			GroupID:         &groupId,
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

func RawToDto(in *group.Group) *proto.Group {
	var members []*proto.User
	for _, usr := range in.Members {
		if usr.IsVerify == nil {
			usr.IsVerify = utils.BoolAdr(false)
		}
		newUser := &proto.User{
			Id:              usr.ID.String(),
			Title:           usr.Title,
			Firstname:       usr.Firstname,
			Lastname:        usr.Lastname,
			Nickname:        usr.Nickname,
			StudentID:       usr.StudentID,
			Faculty:         usr.Faculty,
			Year:            usr.Year,
			Phone:           usr.Phone,
			LineID:          usr.LineID,
			Email:           usr.Email,
			AllergyFood:     usr.AllergyFood,
			FoodRestriction: usr.FoodRestriction,
			AllergyMedicine: usr.AllergyMedicine,
			Disease:         usr.Disease,
			ImageUrl:        "",
			CanSelectBaan:   *usr.CanSelectBaan,
			IsVerify:        *usr.IsVerify,
			GroupId:         usr.GroupID.String(),
		}
		members = append(members, newUser)
	}

	return &proto.Group{
		Id:       in.ID.String(),
		LeaderID: in.LeaderID,
		Token:    in.Token,
		Members:  members,
	}
}
