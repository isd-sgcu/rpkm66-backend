package user

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
)

type User struct {
	model.Base
	Firstname             string `json:"firstname" gorm:"type:tinytext"`
	Lastname              string `json:"lastname" gorm:"type:tinytext"`
	Nickname              string `json:"nickname" gorm:"type:tinytext"`
	StudentID             string `json:"student_id" gorm:"index:,unique"`
	Faculty               string `json:"faculty" gorm:"type:tinytext"`
	Year                  string `json:"year" gorm:"type:tinytext"`
	Phone                 string `json:"phone" gorm:"type:tinytext"`
	LineID                string `json:"line_id" gorm:"type:tinytext"`
	Email                 string `json:"email" gorm:"type:tinytext"`
	AllergyFood           string `json:"allergy_food" gorm:"type:tinytext"`
	FoodRestriction       string `json:"food_restriction" gorm:"type:tinytext"`
	AllergyMedicine       string `json:"allergy_medicine" gorm:"type:tinytext"`
	Disease               string `json:"disease" gorm:"type:tinytext"`
	VaccineCertificateUrl string `json:"vaccine_certificate_url" gorm:"type:varchar(500)"`
	ImageUrl              string `json:"image_url" gorm:"type:varchar(500)"`
}
