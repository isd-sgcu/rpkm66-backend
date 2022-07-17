package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"math/rand"
	"strconv"
	"time"
)

func (s Seed) GroupSeed1657907505114() error {
	for i := 0; i < 10; i++ {
		grp := &group.Group{
			LeaderID: "",
			Token:    strconv.Itoa(int(time.Now().Unix())),
		}
		err := s.db.Create(&grp).Error

		usr := &user.User{
			Title:           faker.Word(),
			Firstname:       faker.FirstName(),
			Lastname:        faker.LastName(),
			Nickname:        faker.Name(),
			StudentID:       strconv.Itoa(int(time.Now().Unix())),
			Faculty:         faker.Word(),
			Year:            faker.Word(),
			Phone:           faker.Phonenumber(),
			LineID:          faker.Word(),
			Email:           faker.Email(),
			AllergyFood:     faker.Word(),
			FoodRestriction: faker.Word(),
			AllergyMedicine: faker.Word(),
			Disease:         faker.Word(),
			CanSelectBaan:   utils.BoolAdr(false),
			GroupID:         &grp.ID,
		}
		err = s.db.Create(&usr).Error
		if err != nil {
			return err
		}

		grp.LeaderID = usr.ID.String()
		err = s.db.Save(&grp).Error
		if err != nil {
			return err
		}

		for i := 0; i < rand.Intn(3); i++ {
			usr := &user.User{
				Title:           faker.Word(),
				Firstname:       faker.FirstName(),
				Lastname:        faker.LastName(),
				Nickname:        faker.Name(),
				StudentID:       strconv.Itoa(int(time.Now().Unix())),
				Faculty:         faker.Word(),
				Year:            faker.Word(),
				Phone:           faker.Phonenumber(),
				LineID:          faker.Word(),
				Email:           faker.Email(),
				AllergyFood:     faker.Word(),
				FoodRestriction: faker.Word(),
				AllergyMedicine: faker.Word(),
				Disease:         faker.Word(),
				CanSelectBaan:   utils.BoolAdr(false),
				GroupID:         &grp.ID,
			}
			err = s.db.Create(&usr).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
