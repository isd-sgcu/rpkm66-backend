package event

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindAllEvent(result *[]*event.Event) error {
	return r.db.Model(&event.Event{}).Find(result).Error
}

func (r *repositoryImpl) FindAllEventWithType(eventType string, result *[]*event.Event) error {
	return r.db.Model(&event.Event{}).Order("events.order ASC").Find(result, "event_type = ?", eventType).Error
}

func (r *repositoryImpl) FindEventByID(id string, result *event.Event) error {
	return r.db.First(&result, "id = ?", id).Error
}

func (r *repositoryImpl) Create(in *event.Event) error {
	return r.db.Create(&in).Error
}

func (r *repositoryImpl) Update(id string, result *event.Event) error {
	return r.db.Where(id, "id = ?", id).Updates(&result).First(&result, "id = ?", id).Error
}

func (r *repositoryImpl) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&event.Event{}).Error
}
