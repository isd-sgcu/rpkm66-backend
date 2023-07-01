package checkin

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/checkin"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) Checkin(ci *checkin.Checkin) error {
	return r.db.Create(ci).Error
}

func (r *repositoryImpl) Checkout(ci *checkin.Checkin) error {
	return r.db.Model(ci).Where("user_id = ? AND event_type = ? AND checkout_at IS NULL", ci.UserId, ci.EventType).Update("checkout_at", ci.CheckoutAt).Error
}

func (r *repositoryImpl) FindLastCheckin(userid string, eventType int32, result *checkin.Checkin) error {
	return r.db.Where("user_id = ? AND event_type = ?", userid, eventType).Order("checkin_at DESC").First(result).Error
}
