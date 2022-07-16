package group

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
)

type Group struct {
	model.Base
	LeaderID string       `json:"leader_id"`
	Token    string       `json:"token" gorm:"index"`
	Members  []*user.User `json:"members"`
}
