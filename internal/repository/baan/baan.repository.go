package baan

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/baan"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindAll(result *[]*baan.Baan) error {
	return r.db.Model(&baan.Baan{}).Find(result).Error
}

func (r *repositoryImpl) FindOne(id string, result *baan.Baan) error {
	return r.db.First(result, "id = ?", id).Error
}

func (r *repositoryImpl) FindMulti(ids []string, result *[]*baan.Baan) error {
	return r.db.Where("id IN ?", ids).Find(&result).Error
}
