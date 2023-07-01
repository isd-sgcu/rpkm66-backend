package group

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/group"
	group_repo "github.com/isd-sgcu/rpkm66-backend/internal/repository/group"
	"gorm.io/gorm"
)

type Repository interface {
	FindGroupByToken(token string, result *group.Group) error
	FindGroupById(id string, result *group.Group) error
	FindGroupWithBaans(id string, result *group.Group) error
	Create(in *group.Group) error
	UpdateWithLeader(leaderId string, result *group.Group) error
	Delete(id string) error
	RemoveAllBaan(g *group.Group) error
}

func NewRepository(db *gorm.DB) Repository {
	return group_repo.NewRepository(db)
}
