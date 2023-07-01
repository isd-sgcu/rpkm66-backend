package checkin

import (
	"context"

	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
	checkin_svc "github.com/isd-sgcu/rpkm66-backend/internal/service/checkin"
	cache_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/cache"
	checkin_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/checkin"
	"github.com/isd-sgcu/rpkm66-backend/proto"
)

type Service interface {
	CheckinVerify(_ context.Context, req *proto.CheckinVerifyRequest) (*proto.CheckinVerifyResponse, error)
	CheckinConfirm(_ context.Context, req *proto.CheckinConfirmRequest) (*proto.CheckinConfirmResponse, error)
}

func NewService(repo checkin_repo.Repository, cache cache_repo.Repository, conf cfgldr.App) Service {
	return checkin_svc.NewService(repo, cache, conf)
}
