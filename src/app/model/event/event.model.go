package event

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
)

type Event struct {
	model.Base
	NameTH        string `json:"name_th"`
	DescriptionTH string `json:"description_th"`
	NameEN        string `json:"name_en"`
	DescriptionEN string `json:"description_en"`
	Code          string `json:"code"`
	ImageURL      string `json:"image_url"`
	EventType     string `json:"event_type"`
	Order         int    `json:"order"`
}
