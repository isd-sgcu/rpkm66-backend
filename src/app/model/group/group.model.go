package group

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"gorm.io/gorm"
)

type Group struct {
	model.Base
	LeaderID string       `json:"leader_id"`
	Token    string       `json:"token" gorm:"index:, unique"`
	Members  []*user.User `json:"members"`
	Baans    []*baan.Baan `json:"baans" gorm:"many2many:group_baans;"`
}

func (m *Group) BeforeCreate(_ *gorm.DB) error {
	m.Token = utils.GenToken(m.LeaderID)
	m.ID = uuid.New()
	return nil
}
