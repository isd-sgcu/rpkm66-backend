package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
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
	FindOne(string, *user.User) error
	FindByStudentID(string, *user.User) error
	Create(*user.User) error
	Update(string, *user.User) error
	Delete(string) error
	CreateOrUpdate(*user.User) error
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.FindOneUserResponse, err error) {
	raw := user.User{}

	err = s.repo.FindOne(req.Id, &raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &proto.FindOneUserResponse{User: RawToDto(&raw)}, nil
}

func (s *Service) FindByStudentID(_ context.Context, req *proto.FindByStudentIDUserRequest) (res *proto.FindByStudentIDUserResponse, err error) {
	result := user.User{}

	err = s.repo.FindByStudentID(req.StudentId, &result)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res = &proto.FindByStudentIDUserResponse{User: RawToDto(&result)}

	return res, err
}

func (s *Service) Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.CreateUserResponse, err error) {
	raw, _ := DtoToRaw(req.User)

	err = s.repo.Create(raw)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &proto.CreateUserResponse{User: RawToDto(raw)}, nil
}

func (s *Service) CreateOrUpdate(_ context.Context, req *proto.CreateOrUpdateUserRequest) (result *proto.CreateOrUpdateUserResponse, err error) {
	raw, err := DtoToRaw(req.User)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	err = s.repo.CreateOrUpdate(raw)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &proto.CreateOrUpdateUserResponse{User: RawToDto(raw)}, nil
}

func (s *Service) Update(_ context.Context, req *proto.UpdateUserRequest) (res *proto.UpdateUserResponse, err error) {
	raw, err := DtoToRaw(req.User)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	err = s.repo.Update(req.User.Id, raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &proto.UpdateUserResponse{User: RawToDto(raw)}, nil
}

func (s *Service) Delete(_ context.Context, req *proto.DeleteUserRequest) (res *proto.DeleteUserResponse, err error) {
	err = s.repo.Delete(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &proto.DeleteUserResponse{Success: true}, nil
}

func DtoToRaw(in *proto.User) (result *user.User, err error) {
	var id uuid.UUID

	if in.Id != "" {
		id, err = uuid.Parse(in.Id)
		if err != nil {
			return nil, err
		}
	}

	return &user.User{
		Base: model.Base{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Title:           in.Title,
		Firstname:       in.Firstname,
		Lastname:        in.Lastname,
		Nickname:        in.Nickname,
		StudentID:       in.StudentID,
		Faculty:         in.Faculty,
		Year:            in.Year,
		Phone:           in.Phone,
		LineID:          in.LineID,
		Email:           in.Email,
		AllergyFood:     in.AllergyFood,
		FoodRestriction: in.FoodRestriction,
		AllergyMedicine: in.AllergyMedicine,
		Disease:         in.Disease,
		ImageUrl:        in.ImageUrl,
		CanSelectBaan:   in.CanSelectBaan,
	}, nil
}
func RawToDto(in *user.User) *proto.User {
	return &proto.User{
		Id:              in.ID.String(),
		Title:           in.Title,
		Firstname:       in.Firstname,
		Lastname:        in.Lastname,
		Nickname:        in.Nickname,
		StudentID:       in.StudentID,
		Faculty:         in.Faculty,
		Year:            in.Year,
		Phone:           in.Phone,
		LineID:          in.LineID,
		Email:           in.Email,
		AllergyFood:     in.AllergyFood,
		FoodRestriction: in.FoodRestriction,
		AllergyMedicine: in.AllergyMedicine,
		Disease:         in.Disease,
		ImageUrl:        in.ImageUrl,
		CanSelectBaan:   in.CanSelectBaan,
	}
}
