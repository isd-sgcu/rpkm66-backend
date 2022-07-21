package baan_group_selection

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	"time"
)

type BaanGroupSelection struct {
	Baan      *baan.Baan
	BaanID    *uuid.UUID `gorm:"primaryKey"`
	GroupID   *uuid.UUID `gorm:"primaryKey"`
	Order     int        `json:"order" gorm:"type:tinyint"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:datetime;autoCreateTime:nano"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:datetime;autoUpdateTime:nano"`
}
