package baan_group_selection

import (
	baan_group_selection "github.com/isd-sgcu/rpkm66-backend/internal/entity/baan-group-selection"
	baan_group_selection_repo "github.com/isd-sgcu/rpkm66-backend/internal/repository/baan-group-selection"
	"gorm.io/gorm"
)

type Repository interface {
	FindBaans(groupId string, result *[]*baan_group_selection.BaanGroupSelection) error
	SaveBaansSelection(result *[]*baan_group_selection.BaanGroupSelection) error
}

func NewRepository(db *gorm.DB) Repository {
	return baan_group_selection_repo.NewRepository(db)
}
