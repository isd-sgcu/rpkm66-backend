package database

import (
	"fmt"
	"strconv"

	"github.com/isd-sgcu/rpkm66-backend/cfgldr"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/baan"
	baan_group "github.com/isd-sgcu/rpkm66-backend/internal/entity/baan-group-selection"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/checkin"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/group"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(conf *cfgldr.Database) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.Host, conf.User, conf.Password, conf.Name, strconv.Itoa(conf.Port))

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	err = db.SetupJoinTable(&group.Group{}, "Baans", &baan_group.BaanGroupSelection{})

	err = db.AutoMigrate(checkin.Checkin{}, group.Group{}, baan.Baan{}, user.User{}, event.Event{})
	if err != nil {
		return nil, err
	}

	return
}
