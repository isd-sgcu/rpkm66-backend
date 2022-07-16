package group

import (
	"context"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	repo IRepository
}

type IRepository interface {
	FindUserById(string, *user.User) error
	FindGroupByToken(string, *group.Group) error
	Create(*group.Group) error
	UpdateUser(*user.User) error
	UpdateWithLeader(string, *group.Group) error
	Delete(string) error
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByToken(_ context.Context, req *proto.FindByTokenGroupRequest) (res *proto.FindByTokenGroupResponse, err error) {
	raw := group.Group{}

	err = s.repo.FindGroupByToken(req.Token, &raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}
	return &proto.FindByTokenGroupResponse{Group: RawToDto(&raw)}, nil
}

func (s *Service) Create(_ context.Context, req *proto.CreateGroupRequest) (res *proto.CreateGroupResponse, err error) {
	if _, err = uuid.Parse(req.UserId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid leader id")
	}

	usr := &user.User{}
	err = s.repo.FindUserById(req.UserId, usr)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	in := &group.Group{
		LeaderID: req.UserId,
	}
	err = s.repo.Create(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create group")
	}

	usr.GroupID = &in.ID
	err = s.repo.UpdateUser(usr)
	in.Members = []*user.User{usr}

	return &proto.CreateGroupResponse{Group: RawToDto(in)}, nil
}

func (s *Service) Update(_ context.Context, req *proto.UpdateGroupRequest) (res *proto.UpdateGroupResponse, err error) {
	raw, err := DtoToRaw(req.Group)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid group id")
	}

	err = s.repo.UpdateWithLeader(req.LeaderId, raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}

	return &proto.UpdateGroupResponse{Group: RawToDto(raw)}, nil
}

func (s *Service) Join(_ context.Context, req *proto.JoinGroupRequest) (res *proto.JoinGroupResponse, err error) {
	if req.IsLeader && req.Members > 1 {
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	if _, err = uuid.Parse(req.UserId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	joinGroup := &group.Group{}
	err = s.repo.FindGroupByToken(req.Token, joinGroup)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}

	if len(joinGroup.Members) >= 3 {
		return nil, status.Error(codes.PermissionDenied, "group full")
	}

	joinUser := &user.User{}
	err = s.repo.FindUserById(req.UserId, joinUser)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	prevGroupId := joinUser.GroupID
	joinUser.GroupID = &joinGroup.ID
	err = s.repo.UpdateUser(joinUser)

	if req.IsLeader && req.Members == 1 {
		err = s.repo.Delete(prevGroupId.String())
	}

	joinGroup.Members = append(joinGroup.Members, joinUser)
	return &proto.JoinGroupResponse{Group: RawToDto(joinGroup)}, nil
}

func (s *Service) DeleteMember(_ context.Context, req *proto.DeleteMemberGroupRequest) (res *proto.DeleteMemberGroupResponse, err error) {
	_, err = uuid.Parse(req.DeletedId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	deletedGrp := &group.Group{}
	err = s.repo.FindGroupByToken(req.Token, deletedGrp)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}

	if deletedGrp.LeaderID != req.UserId {
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	//create a new group for removed user
	newGroup := &group.Group{
		LeaderID: req.DeletedId,
	}
	err = s.repo.Create(newGroup)

	//remove the user from the group
	var removedMember []*user.User
	var removedUser *user.User
	for _, usr := range deletedGrp.Members {
		if usr.ID.String() == req.DeletedId {
			removedUser = usr
		} else {
			removedMember = append(removedMember, usr)
		}
	}
	deletedGrp.Members = removedMember
	removedUser.GroupID = &newGroup.ID
	err = s.repo.UpdateUser(removedUser)

	return &proto.DeleteMemberGroupResponse{Group: RawToDto(deletedGrp)}, nil
}

func (s *Service) Leave(_ context.Context, req *proto.LeaveGroupRequest) (res *proto.LeaveGroupResponse, err error) {
	if _, err = uuid.Parse(req.UserId); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	prevGrp := &group.Group{}
	err = s.repo.FindGroupByToken(req.Token, prevGrp)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}

	if req.UserId == prevGrp.LeaderID {
		return nil, status.Error(codes.PermissionDenied, "not allowed")
	}

	leavedUser := &user.User{}
	for _, usr := range prevGrp.Members {
		if usr.ID.String() == req.UserId {
			leavedUser = usr
			break
		}
	}
	in := &group.Group{
		LeaderID: req.UserId,
	}
	err = s.repo.Create(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create group")
	}

	leavedUser.GroupID = &in.ID
	err = s.repo.UpdateUser(leavedUser)
	in.Members = []*user.User{leavedUser}
	return &proto.LeaveGroupResponse{Group: RawToDto(in)}, nil
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
