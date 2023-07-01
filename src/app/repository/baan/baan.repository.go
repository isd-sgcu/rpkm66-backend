package baan

import (
	"github.com/isd-sgcu/rpkm66-backend/src/app/entity/baan"
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

func (r *Repository) FindMulti(ids []string, result *[]*baan.Baan) error {
	return r.db.Where("id IN ?", ids).Find(&result).Error
}
