package event

import (
	entity "github.com/isd-sgcu/rpkm66-backend/internal/entity"
)

type Event struct {
	entity.Base
	NameTH        string `json:"name_th"`
	DescriptionTH string `json:"description_th"`
	NameEN        string `json:"name_en"`
	DescriptionEN string `json:"description_en"`
	Code          string `json:"code"`
	ImageURL      string `json:"image_url"`
	EventType     string `json:"event_type"`
	Order         int    `json:"order"`
}
