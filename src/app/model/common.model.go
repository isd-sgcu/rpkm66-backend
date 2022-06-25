package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime;autoCreateTime:nano"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime;autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;type:datetime"`
}

func (b *Base) BeforeCreate(_ *gorm.DB) error {
	b.ID = uuid.New()

	return nil
}
