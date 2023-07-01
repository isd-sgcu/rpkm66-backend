package checkin

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/checkin"
	checkin_repo "github.com/isd-sgcu/rpkm66-backend/internal/repository/checkin"
	"gorm.io/gorm"
)

type Repository interface {
	Checkin(ci *checkin.Checkin) error
	Checkout(ci *checkin.Checkin) error
	FindLastCheckin(userid string, eventType int32, result *checkin.Checkin) error
}

func NewRepository(db *gorm.DB) Repository {
	return checkin_repo.NewRepository(db)
}
