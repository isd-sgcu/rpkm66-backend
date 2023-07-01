package group

import (
	"context"

	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/group"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/user"
	group_svc "github.com/isd-sgcu/rpkm66-backend/internal/service/group"
	baan_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/baan"
	baan_group_selection_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/baan-group-selection"
	cache_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/cache"
	group_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/group"
	user_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/user"
	file_svc "github.com/isd-sgcu/rpkm66-backend/pkg/service/file"
	"github.com/isd-sgcu/rpkm66-backend/proto"
)

type Service interface {
	FindOne(_ context.Context, req *proto.FindOneGroupRequest) (res *proto.FindOneGroupResponse, err error)
	FindByToken(_ context.Context, req *proto.FindByTokenGroupRequest) (res *proto.FindByTokenGroupResponse, err error)
	Update(_ context.Context, req *proto.UpdateGroupRequest) (res *proto.UpdateGroupResponse, err error)
	Join(_ context.Context, req *proto.JoinGroupRequest) (res *proto.JoinGroupResponse, err error)
	DeleteMember(_ context.Context, req *proto.DeleteMemberGroupRequest) (res *proto.DeleteMemberGroupResponse, err error)
	Leave(_ context.Context, req *proto.LeaveGroupRequest) (res *proto.LeaveGroupResponse, err error)
	SelectBaan(_ context.Context, req *proto.SelectBaanRequest) (res *proto.SelectBaanResponse, err error)
	RawToDto(in *group.Group) (result *proto.Group, err error)
	GetMembersImages(users []*user.User) (result []*proto.UserInfo, err error)
	GetUserImage(usr *user.User) (result string, err error)
}

func NewService(repo group_repo.Repository, userRepo user_repo.Repository, baanGroupSelectionRepo baan_group_selection_repo.Repository, fileSrv file_svc.Service, cacheRepo cache_repo.Repository, baanRepo baan_repo.Repository, conf cfgldr.App) Service {
	return group_svc.NewService(repo, userRepo, baanGroupSelectionRepo, fileSrv, cacheRepo, baanRepo, conf)
}
