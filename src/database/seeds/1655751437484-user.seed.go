package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
)

func (s Seed) UserSeed1655751437484() error {
	for i := 0; i < 10; i++ {
		usr := user.User{
			Title:           faker.Word(),
			Firstname:       faker.FirstName(),
			Lastname:        faker.LastName(),
			Nickname:        faker.Name(),
			StudentID:       faker.Word(),
			Faculty:         faker.Word(),
			Year:            faker.Word(),
			Phone:           faker.Phonenumber(),
			LineID:          faker.Word(),
			Email:           faker.Email(),
			AllergyFood:     faker.Word(),
			FoodRestriction: faker.Word(),
			AllergyMedicine: faker.Word(),
			Disease:         faker.Word(),
			ImageUrl:        faker.URL(),
		}
		err := s.db.Create(&usr).Error

		if err != nil {
			return err
		}
	}
	return nil
}
