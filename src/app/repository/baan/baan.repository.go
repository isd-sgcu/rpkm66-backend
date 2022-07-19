package baan

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(result *[]*baan.Baan) error {
	return r.db.Model(&baan.Baan{}).Find(result).Error
}

func (r *Repository) FindOne(id string, result *baan.Baan) error {
	return r.db.First(result, "id = ?", id).Error
}
