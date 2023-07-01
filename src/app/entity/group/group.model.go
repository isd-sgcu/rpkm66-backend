package group

import (
	"github.com/google/uuid"
	entity "github.com/isd-sgcu/rpkm66-backend/src/app/entity"
	"github.com/isd-sgcu/rpkm66-backend/src/app/entity/baan"
	baan_group_selection "github.com/isd-sgcu/rpkm66-backend/src/app/entity/baan-group-selection"
	"github.com/isd-sgcu/rpkm66-backend/src/app/entity/user"
	"github.com/isd-sgcu/rpkm66-backend/src/app/utils"
	"gorm.io/gorm"
)

type Group struct {
	entity.Base
	LeaderID           string       `json:"leader_id"`
	Token              string       `json:"token" gorm:"index:, unique"`
	Members            []*user.User `json:"members"`
	Baans              []*baan.Baan `json:"baans" gorm:"many2many:group_baans;"`
	BaanGroupSelection []*baan_group_selection.BaanGroupSelection
}

func (m *Group) BeforeCreate(_ *gorm.DB) error {
	m.Token = utils.GenToken(m.LeaderID)
	m.ID = uuid.New()
	return nil
}
