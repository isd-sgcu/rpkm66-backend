package group

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/group"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindGroupByToken(token string, result *group.Group) error {
	return r.db.Preload("Members").First(&result, "token = ?", token).Error
}

func (r *repositoryImpl) FindGroupById(id string, result *group.Group) error {
	return r.db.Preload("Members").First(&result, "id = ?", id).Error
}

func (r *repositoryImpl) FindGroupWithBaans(id string, result *group.Group) error {
	return r.db.Preload("BaanGroupSelection", func(db *gorm.DB) *gorm.DB {
		return db.Order("baan_group_selections.order ASC")
	}).Preload("BaanGroupSelection.Baan").Where("id = ?", id).Find(&result).Error
}

func (r *repositoryImpl) Create(in *group.Group) error {
	return r.db.Create(&in).Error
}

func (r *repositoryImpl) UpdateWithLeader(leaderId string, result *group.Group) error {
	return r.db.Where("leader_id = ?", leaderId).Updates(&result).First(&result, "leader_id = ?", leaderId).Error
}

func (r *repositoryImpl) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&group.Group{}).Error
}

func (r *repositoryImpl) RemoveAllBaan(g *group.Group) error {
	return r.db.Model(&g).Association("Baans").Clear()
}
