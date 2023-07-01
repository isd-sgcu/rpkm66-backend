package group

import (
	"github.com/isd-sgcu/rpkm66-backend/src/app/entity/group"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindGroupByToken(token string, result *group.Group) error {
	return r.db.Preload("Members").First(&result, "token = ?", token).Error
}

func (r *Repository) FindGroupById(id string, result *group.Group) error {
	return r.db.Preload("Members").First(&result, "id = ?", id).Error
}

func (r *Repository) FindGroupWithBaans(id string, result *group.Group) error {
	return r.db.Preload("BaanGroupSelection", func(db *gorm.DB) *gorm.DB {
		return db.Order("baan_group_selections.order ASC")
	}).Preload("BaanGroupSelection.Baan").Where("id = ?", id).Find(&result).Error
}

func (r *Repository) Create(in *group.Group) error {
	return r.db.Create(&in).Error
}

func (r *Repository) UpdateWithLeader(leaderId string, result *group.Group) error {
	return r.db.Where("leader_id = ?", leaderId).Updates(&result).First(&result, "leader_id = ?", leaderId).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&group.Group{}).Error
}

func (r *Repository) RemoveAllBaan(g *group.Group) error {
	return r.db.Model(&g).Association("Baans").Clear()
}
