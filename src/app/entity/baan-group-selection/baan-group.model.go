package baan_group_selection

import (
	"time"

	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm66-backend/src/app/entity/baan"
)

type BaanGroupSelection struct {
	Baan      *baan.Baan
	BaanID    *uuid.UUID `gorm:"primaryKey"`
	GroupID   *uuid.UUID `gorm:"primaryKey"`
	Order     int        `json:"order" gorm:"type:tinyint"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp;autoUptimestamp:nano"`
}
