package baan

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/baan"
	baan_repo "github.com/isd-sgcu/rpkm66-backend/internal/repository/baan"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(result *[]*baan.Baan) error
	FindOne(id string, result *baan.Baan) error
	FindMulti(ids []string, result *[]*baan.Baan) error
}

func NewRepository(db *gorm.DB) Repository {
	return baan_repo.NewRepository(db)
}
