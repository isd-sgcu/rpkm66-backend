package checkin

import (
	"time"

	entity "github.com/isd-sgcu/rpkm66-backend/internal/entity"
)

type Checkin struct {
	entity.Base
	UserId     string     `json:"user_id" gorm:"index:flci_idx,priority:1,type:btree;index:co_idx,type:btree,priority:1"`
	CheckinAt  *time.Time `json:"check_in_at" gorm:"type:timestamp;autoCreateTime:nano;index:flci_idx,priority:3,sort:desc"`
	CheckoutAt *time.Time `json:"check_out_at" gorm:"type:timestamp;index:co_idx,priority:3"`
	EventType  int32      `json:"event_type" gorm:"index:flci_idx,priority:2;index:co_idx,priority:2"`
}
