package user

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindOne(id string, result *user.User) error {
	return r.db.First(&result, "id = ?", id).Error
}

func (r *Repository) FindByStudentID(sid string, result *user.User) error {
	return r.db.First(&result, "student_id = ?", sid).Error
}

func (r *Repository) CreateOrUpdate(result *user.User) error {
	if r.db.Where("id = ?", result.ID.String()).Updates(&result).RowsAffected == 0 {
		return r.db.Create(&result).Error
	}
	return nil
}

func (r *Repository) Verify(studentId string) error {
	return r.db.Model(&user.User{}).First(&user.User{}, "student_id = ?", studentId).Update("is_verify", true).Error
}

func (r *Repository) Create(in *user.User) error {
	return r.db.Create(&in).Error
}

func (r *Repository) Update(id string, result *user.User) error {
	return r.db.Where(id, "id = ?", id).Updates(&result).First(&result, "id = ?", id).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.First(id).Delete(&user.User{}).Error
}
