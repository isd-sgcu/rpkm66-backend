package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
)

func (s Seed) EventSeed1658584832819() error {
	for i := 0; i < 10; i++ {
		event := event.Event{
			NameTH:        faker.Word(),
			DescriptionTH: faker.Sentence(),
			NameEN:        faker.Word(),
			DescriptionEN: faker.Sentence(),
			Code:          faker.Word(),
			ImageURL:      faker.URL(),
			EventType:     "estamp",
		}
		err := s.db.Create(&event).Error

		if err != nil {
			return err
		}
	}
	return nil
}
