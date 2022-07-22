package baan_group_selection

import (
	baan_group_selection "github.com/isd-sgcu/rnkm65-backend/src/app/model/baan-group-selection"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveBaansSelection(result *[]*baan_group_selection.BaanGroupSelection) error {
	return r.db.Save(result).Error
}

func (r *Repository) FindBaans(groupId string, result *[]*baan_group_selection.BaanGroupSelection) error {
	return r.db.Model(&baan_group_selection.BaanGroupSelection{}).Order("baan_group_selections.order ASC").Preload("Baan").Find(result, "group_id = ?", groupId).Error
}
