package checkin

import (
	"time"

	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
)

type Checkin struct {
	model.Base
	UserId     string     `json:"user_id" gorm:"index:flci_idx,priority:1,type:btree;index:co_idx,type:btree,priority:1"`
	CheckinAt  *time.Time `json:"check_in_at" gorm:"type:datetime;autoCreateTime:nano;index:flci_idx,priority:3,sort:desc"`
	CheckoutAt *time.Time `json:"check_out_at" gorm:"type:datetime;index:co_idx,priority:3"`
	EventType  int32      `json:"event_type" gorm:"index:flci_idx,priority:2;index:co_idx,priority:2"`
}
