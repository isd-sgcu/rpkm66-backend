package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"strconv"
	"time"
)

func (s Seed) GroupSeed1657907505114() error {
	for i := 0; i < 10; i++ {
		g := group.Group{
			Token: faker.Word() + strconv.Itoa(int(time.Now().Unix())) + faker.Timestamp(),
		}

		err := s.db.Save(&g).Error
		if err != nil {
			return err
		}
	}
	return nil
}
