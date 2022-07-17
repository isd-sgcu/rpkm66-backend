package group

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
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

func (r *Repository) Create(in *group.Group) error {
	in.Token = utils.GenToken(in.LeaderID)
	return r.db.Preload("Members").Create(&in).Error
}

func (r *Repository) Update(result *group.Group) error {
	return r.db.Save(&result).Error
}

func (r *Repository) UpdateWithLeader(leaderId string, result *group.Group) error {
	return r.db.Where("leader_id = ?", leaderId).Updates(&result).First(&result, "leader_id = ?", leaderId).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&group.Group{}).Error
}
