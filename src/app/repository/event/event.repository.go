package event

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAllEvent(result *[]*event.Event) error {
	return r.db.Model(&event.Event{}).Find(result).Error
}

func (r *Repository) FindAllEventWithType(eventType string, result *[]*event.Event) error {
	return r.db.Model(&event.Event{}).Find(result, "event_type = ?", eventType).Error
}

func (r *Repository) FindEventByID(id string, result *event.Event) error {
	return r.db.First(&result, "id = ?", id).Error
}

func (r *Repository) Create(in *event.Event) error {
	return r.db.Create(&in).Error
}

func (r *Repository) Update(id string, result *event.Event) error {
	return r.db.Where(id, "id = ?", id).Updates(&result).First(&result, "id = ?", id).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&event.Event{}).Error
}
