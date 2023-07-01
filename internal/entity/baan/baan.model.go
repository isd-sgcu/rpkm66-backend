package baan

import (
	"github.com/isd-sgcu/rpkm66-backend/constant/baan"
	entity "github.com/isd-sgcu/rpkm66-backend/internal/entity"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/user"
)

type Baan struct {
	entity.Base
	NameTH        string        `json:"name_th" gorm:"type:tinytext"`
	DescriptionTH string        `json:"description_th" gorm:"type:mediumtext"`
	NameEN        string        `json:"name_en" gorm:"type:tinytext"`
	DescriptionEN string        `json:"description_en" gorm:"type:mediumtext"`
	ImageUrl      string        `json:"image_url" gorm:"type:text"`
	Size          baan.BaanSize `json:"size" gorm:"type:tinyint"`
	Facebook      string        `json:"facebook" gorm:"type:tinytext"`
	FacebookUrl   string        `json:"facebook_url" gorm:"type:text"`
	Instagram     string        `json:"instagram" gorm:"type:tinytext"`
	InstagramUrl  string        `json:"instagram_url" gorm:"type:text"`
	Line          string        `json:"line" gorm:"type:tinytext"`
	LineUrl       string        `json:"line_url" gorm:"type:text"`
	Members       []*user.User  `json:"members"`
}
