package checkin

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Checkin(ci *checkin.Checkin) error {
	return r.db.Create(ci).Error
}

func (r *Repository) Checkout(ci *checkin.Checkin) error {
	return r.db.Model(ci).Where("user_id = ? AND event_type = ? AND checkout_at IS NULL", ci.UserId, ci.EventType).Update("checkout_at", ci.CheckoutAt).Error
}

func (r *Repository) FindLastCheckin(userid string, eventType int32, result *checkin.Checkin) error {
	return r.db.Where("user_id = ? AND event_type = ?", userid, eventType).Order("checkin_at DESC").First(result).Error
}
