package baan_group_selection

import (
	baan_group_selection "github.com/isd-sgcu/rpkm66-backend/internal/entity/baan-group-selection"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) SaveBaansSelection(result *[]*baan_group_selection.BaanGroupSelection) error {
	return r.db.Save(result).Error
}

func (r *repositoryImpl) FindBaans(groupId string, result *[]*baan_group_selection.BaanGroupSelection) error {
	return r.db.Model(&baan_group_selection.BaanGroupSelection{}).Order("baan_group_selections.order ASC").Preload("Baan").Find(result, "group_id = ?", groupId).Error
}
