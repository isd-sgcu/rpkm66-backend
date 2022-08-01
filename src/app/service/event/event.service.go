package event

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Service struct {
	repo IRepository
}

type IRepository interface {
	FindAllEvent(*[]*event.Event) error
	FindEventByID(string, *event.Event) error
	Create(*event.Event) error
	Update(string, *event.Event) error
	Delete(string) error
	FindAllEventWithType(string, *[]*event.Event) error
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindAllEvent(_ context.Context, req *proto.FindAllEventRequest) (res *proto.FindAllEventResponse, err error) {
	var events []*event.Event

	err = s.repo.FindAllEvent(&events)
	if err != nil {

		log.Error().Err(err).
			Str("service", "event").
			Str("module", "find all").
			Msg("Error while querying all events")

		return nil, status.Error(codes.Unavailable, "Internal error")
	}

	return &proto.FindAllEventResponse{Event: RawToDtoList(&events)}, nil
}

func (s *Service) FindEventByID(_ context.Context, req *proto.FindEventByIDRequest) (res *proto.FindEventByIDResponse, err error) {
	raw := event.Event{}

	err = s.repo.FindEventByID(req.Id, &raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "event not found")
	}

	return &proto.FindEventByIDResponse{Event: RawToDto(&raw)}, nil
}

func (s *Service) Create(_ context.Context, req *proto.CreateEventRequest) (res *proto.CreateEventResponse, err error) {
	raw, _ := DtoToRaw(req.Event)

	err = s.repo.Create(raw)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create event")
	}

	return &proto.CreateEventResponse{Event: RawToDto(raw)}, nil
}

func (s *Service) Update(_ context.Context, req *proto.UpdateEventRequest) (res *proto.UpdateEventResponse, err error) {
	raw, err := DtoToRaw(req.Event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid event id")
	}

	err = s.repo.Update(req.Event.Id, raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "event not found")
	}

	return &proto.UpdateEventResponse{Event: RawToDto(raw)}, nil
}

func (s *Service) Delete(_ context.Context, req *proto.DeleteEventRequest) (res *proto.DeleteEventResponse, err error) {
	err = s.repo.Delete(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "event not found")
	}

	return &proto.DeleteEventResponse{Success: true}, nil
}

func (s *Service) FindAllEventWithType(_ context.Context, req *proto.FindAllEventWithTypeRequest) (res *proto.FindAllEventWithTypeResponse, err error) {
	var events []*event.Event
	err = s.repo.FindAllEventWithType(req.EventType, &events)
	if err != nil {
		return nil, status.Error(codes.NotFound, "eventType not found")
	}

	return &proto.FindAllEventWithTypeResponse{Event: RawToDtoList(&events)}, nil
}

func DtoToRaw(in *proto.Event) (result *event.Event, err error) {
	var id uuid.UUID

	if in.Id != "" {
		id, err = uuid.Parse(in.Id)
		if err != nil {
			return nil, err
		}
	}

	return &event.Event{
		Base: model.Base{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		NameTH:        in.NameTH,
		DescriptionTH: in.DescriptionTH,
		NameEN:        in.NameEN,
		DescriptionEN: in.DescriptionEN,
		Code:          in.Code,
		ImageURL:      in.ImageURL,
	}, nil
}

func RawToDto(in *event.Event) *proto.Event {
	return &proto.Event{
		Id:            in.ID.String(),
		NameTH:        in.NameTH,
		DescriptionTH: in.DescriptionTH,
		NameEN:        in.NameEN,
		DescriptionEN: in.DescriptionEN,
		Code:          in.Code,
		ImageURL:      in.ImageURL,
	}
}

func RawToDtoList(in *[]*event.Event) []*proto.Event {
	var result []*proto.Event
	for _, e := range *in {
		result = append(result, RawToDto(e))
	}

	return result
}
