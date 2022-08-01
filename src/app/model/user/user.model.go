package user

import (
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"gorm.io/gorm"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
)

type User struct {
	model.Base
	Checkin         []*checkin.Checkin
	Title           string     `json:"title" gorm:"type:VARCHAR(10)"`
	Firstname       string     `json:"firstname" gorm:"type:tinytext"`
	Lastname        string     `json:"lastname" gorm:"type:tinytext"`
	Nickname        string     `json:"nickname" gorm:"type:tinytext"`
	StudentID       string     `json:"student_id" gorm:"index:,unique"`
	Faculty         string     `json:"faculty" gorm:"type:tinytext"`
	Year            string     `json:"year" gorm:"type:tinytext"`
	Phone           string     `json:"phone" gorm:"type:tinytext"`
	LineID          string     `json:"line_id" gorm:"type:tinytext"`
	Email           string     `json:"email" gorm:"type:tinytext"`
	AllergyFood     string     `json:"allergy_food" gorm:"type:tinytext"`
	FoodRestriction string     `json:"food_restriction" gorm:"type:tinytext"`
	AllergyMedicine string     `json:"allergy_medicine" gorm:"type:tinytext"`
	Disease         string     `json:"disease" gorm:"type:tinytext"`
	CanSelectBaan   *bool      `json:"can_select_baan"`
	IsVerify        *bool      `json:"is_verify"`
	GroupID         *uuid.UUID `json:"group_id" gorm:"index"`
	Events          []*event.Event `json:"events" gorm:"many2many:event_user"`
	BaanID          *uuid.UUID `json:"baan_id" gorm:"index"`
}

func (m *User) BeforeCreate(_ *gorm.DB) error {
	m.GroupID = nil
	m.BaanID = nil
	m.ID = uuid.New()

	return nil
}
