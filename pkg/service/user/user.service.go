package user

import (
	"context"

	user_svc "github.com/isd-sgcu/rpkm66-backend/internal/service/user"
	event_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/event"
	user_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/user"
	file_svc "github.com/isd-sgcu/rpkm66-backend/pkg/service/file"
	"github.com/isd-sgcu/rpkm66-backend/proto"
)

type Service interface {
	FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.FindOneUserResponse, err error)
	FindByStudentID(_ context.Context, req *proto.FindByStudentIDUserRequest) (res *proto.FindByStudentIDUserResponse, err error)
	Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.CreateUserResponse, err error)
	CreateOrUpdate(_ context.Context, req *proto.CreateOrUpdateUserRequest) (result *proto.CreateOrUpdateUserResponse, err error)
	Verify(_ context.Context, req *proto.VerifyUserRequest) (res *proto.VerifyUserResponse, err error)
	Update(_ context.Context, req *proto.UpdateUserRequest) (res *proto.UpdateUserResponse, err error)
	Delete(_ context.Context, req *proto.DeleteUserRequest) (res *proto.DeleteUserResponse, err error)
	ConfirmEstamp(_ context.Context, req *proto.ConfirmEstampRequest) (res *proto.ConfirmEstampResponse, err error)
	GetUserEstamp(_ context.Context, req *proto.GetUserEstampRequest) (res *proto.GetUserEstampResponse, err error)
}

func NewService(repo user_repo.Repository, fileSrv file_svc.Service, eventRepo event_repo.Repository) Service {
	return user_svc.NewService(repo, fileSrv, eventRepo)
}
