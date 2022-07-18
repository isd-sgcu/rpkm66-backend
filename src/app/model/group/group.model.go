package group

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"gorm.io/gorm"
)

type Group struct {
	model.Base
	LeaderID string       `json:"leader_id"`
	Token    string       `json:"token" gorm:"index:, unique"`
	Members  []*user.User `json:"members"`
}

func (u *Group) BeforeCreate(_ *gorm.DB) error {
	u.Token = utils.GenToken(u.LeaderID)
	u.ID = uuid.New()
	return nil
}
