package baan

import (
	"context"

	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
	proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/baan/v1"
	baan_svc "github.com/isd-sgcu/rpkm66-backend/internal/service/baan"
	baan_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/baan"
	cache_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/cache"
)

type Service interface {
	FindAllBaan(_ context.Context, _ *proto.FindAllBaanRequest) (*proto.FindAllBaanResponse, error)
	FindOneBaan(_ context.Context, req *proto.FindOneBaanRequest) (*proto.FindOneBaanResponse, error)
}

func NewService(repository baan_repo.Repository, cache cache_repo.Repository, conf cfgldr.App) Service {
	return baan_svc.NewService(repository, cache, conf)
}
