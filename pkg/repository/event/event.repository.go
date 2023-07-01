package event

import (
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	event_repo "github.com/isd-sgcu/rpkm66-backend/internal/repository/event"
	"gorm.io/gorm"
)

type Repository interface {
	FindAllEvent(result *[]*event.Event) error
	FindAllEventWithType(eventType string, result *[]*event.Event) error
	FindEventByID(id string, result *event.Event) error
	Create(in *event.Event) error
	Update(id string, result *event.Event) error
	Delete(id string) error
}

func NewRepository(db *gorm.DB) Repository {
	return event_repo.NewRepository(db)
}
