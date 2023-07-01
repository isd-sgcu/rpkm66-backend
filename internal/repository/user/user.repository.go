package user

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/user"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) FindOne(id string, result *user.User) error {
	return r.db.First(&result, "id = ?", id).Error
}

func (r *repositoryImpl) FindByStudentID(sid string, result *user.User) error {
	return r.db.First(&result, "student_id = ?", sid).Error
}

func (r *repositoryImpl) CreateOrUpdate(result *user.User) error {
	if r.db.Where("id = ?", result.ID.String()).Updates(&result).RowsAffected == 0 {
		return r.db.Create(&result).Error
	}
	return nil
}

func (r *repositoryImpl) Verify(studentId string, columnName string) error {
	return r.db.Model(&user.User{}).First(&user.User{}, "student_id = ?", studentId).Update(columnName, true).Error
}

func (r *repositoryImpl) Create(in *user.User) error {
	return r.db.Create(&in).Error
}

func (r *repositoryImpl) Update(id string, result *user.User) error {
	return r.db.Where(id, "id = ?", id).Updates(&result).First(&result, "id = ?", id).Error
}

func (r *repositoryImpl) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&user.User{}).Error
}

func (r *repositoryImpl) ConfirmEstamp(uId string, thisUser *user.User, thisEvent *event.Event) error {
	//Add this estamp to user
	return r.db.Model(thisUser).Where("id = ?", uId).Omit("Events.*").Association("Events").Append(thisEvent)
}

func (r *repositoryImpl) GetUserEstamp(uId string, thisUser *user.User, results *[]*event.Event) error {
	//Get all estamp that this user has
	return r.db.Model(thisUser).Where("user_id", uId).Association("Events").Find(&results)
}
