package seed

import (
	"strconv"
	"time"

	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"

	"github.com/bxcodec/faker/v3"
)

func (s Seed) UserSeed1655751437484() error {
	for i := 0; i < 10; i++ {
		usr := user.User{
			Title:           getTitle(),
			Firstname:       faker.FirstName(),
			Lastname:        faker.LastName(),
			Nickname:        faker.Name(),
			StudentID:       faker.Word() + strconv.Itoa(int(time.Now().Unix())),
			Faculty:         faker.Word(),
			Year:            faker.Word(),
			Phone:           faker.Phonenumber(),
			LineID:          faker.Word(),
			Email:           faker.Email(),
			AllergyFood:     faker.Word(),
			FoodRestriction: faker.Word(),
			AllergyMedicine: faker.Word(),
			Disease:         faker.Word(),
		}
		err := s.db.Create(&usr).Error

		if err != nil {
			return err
		}
	}
	return nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getTitle() string {
	title := faker.Word()
	return title[0:min(10, len(title))]
}
