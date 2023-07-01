package event

import (
	"context"

	event_svc "github.com/isd-sgcu/rpkm66-backend/internal/service/event"
	event_repo "github.com/isd-sgcu/rpkm66-backend/pkg/repository/event"
	"github.com/isd-sgcu/rpkm66-backend/proto"
)

type Service interface {
	FindAllEvent(_ context.Context, req *proto.FindAllEventRequest) (res *proto.FindAllEventResponse, err error)
	FindEventByID(_ context.Context, req *proto.FindEventByIDRequest) (res *proto.FindEventByIDResponse, err error)
	Create(_ context.Context, req *proto.CreateEventRequest) (res *proto.CreateEventResponse, err error)
	Update(_ context.Context, req *proto.UpdateEventRequest) (res *proto.UpdateEventResponse, err error)
	Delete(_ context.Context, req *proto.DeleteEventRequest) (res *proto.DeleteEventResponse, err error)
	FindAllEventWithType(_ context.Context, req *proto.FindAllEventWithTypeRequest) (res *proto.FindAllEventWithTypeResponse, err error)
}

func NewService(repo event_repo.Repository) Service {
	return event_svc.NewService(repo)
}
