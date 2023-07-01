package user

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/user"
	user_repo "github.com/isd-sgcu/rpkm66-backend/internal/repository/user"
	"gorm.io/gorm"
)

type Repository interface {
	FindOne(id string, result *user.User) error
	FindByStudentID(sid string, result *user.User) error
	CreateOrUpdate(result *user.User) error
	Verify(studentId string, columnName string) error
	Create(in *user.User) error
	Update(id string, result *user.User) error
	Delete(id string) error
	ConfirmEstamp(uId string, thisUser *user.User, thisEvent *event.Event) error
	GetUserEstamp(uId string, thisUser *user.User, results *[]*event.Event) error
}

func NewRepository(db *gorm.DB) Repository {
	return user_repo.NewRepository(db)
}
