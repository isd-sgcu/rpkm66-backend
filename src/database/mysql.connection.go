package database

import (
	"fmt"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/baan"
	baan_group "github.com/isd-sgcu/rpkm66-backend/src/app/model/baan-group-selection"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/checkin"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/event"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/group"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/user"
	"github.com/isd-sgcu/rpkm66-backend/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func InitDatabase(conf *config.Database) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", conf.User, conf.Password, conf.Host, strconv.Itoa(conf.Port), conf.Name)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.SetupJoinTable(&group.Group{}, "Baans", &baan_group.BaanGroupSelection{})

	err = db.AutoMigrate(checkin.Checkin{}, group.Group{}, baan.Baan{}, user.User{}, event.Event{})
	if err != nil {
		return nil, err
	}

	return
}
