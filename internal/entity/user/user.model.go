package user

import (
	"github.com/google/uuid"
	entity "github.com/isd-sgcu/rpkm66-backend/internal/entity"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/checkin"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	"gorm.io/gorm"
)

type User struct {
	entity.Base
	Checkin         []*checkin.Checkin
	Title           string         `json:"title" gorm:"type:VARCHAR(10)"`
	Firstname       string         `json:"firstname" gorm:"type:text"`
	Lastname        string         `json:"lastname" gorm:"type:text"`
	Nickname        string         `json:"nickname" gorm:"type:text"`
	StudentID       string         `json:"student_id" gorm:"index:,unique"`
	Faculty         string         `json:"faculty" gorm:"type:text"`
	Year            string         `json:"year" gorm:"type:text"`
	Phone           string         `json:"phone" gorm:"type:text"`
	LineID          string         `json:"line_id" gorm:"type:text"`
	Email           string         `json:"email" gorm:"type:text"`
	AllergyFood     string         `json:"allergy_food" gorm:"type:text"`
	FoodRestriction string         `json:"food_restriction" gorm:"type:text"`
	AllergyMedicine string         `json:"allergy_medicine" gorm:"type:text"`
	Disease         string         `json:"disease" gorm:"type:text"`
	EmerPhone       string         `json:"emer_phone" gorm:"type:text"`
	EmerRelation    string         `json:"emer_relation" gorm:"type:text"`
	WantBottle      *bool          `json:"wants_bottle" gorm:"type:boolean"`
	CanSelectBaan   *bool          `json:"can_select_baan"`
	IsVerify        *bool          `json:"is_verify"`
	IsGotTicket     *bool          `json:"is_got_ticket"`
	GroupID         *uuid.UUID     `json:"group_id" gorm:"index"`
	Events          []*event.Event `json:"events" gorm:"many2many:event_user"`
	BaanID          *uuid.UUID     `json:"baan_id" gorm:"index"`
	PersonalityGame string         `json:"personality_game" gorm:"type:text"`
}

func (m *User) BeforeCreate(_ *gorm.DB) error {
	m.GroupID = nil
	m.BaanID = nil
	m.ID = uuid.New()

	return nil
}
