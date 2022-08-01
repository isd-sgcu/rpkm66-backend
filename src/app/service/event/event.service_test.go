package event

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/event"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type EventServiceTest struct {
	suite.Suite
	Event              *event.Event
	EventDto           *proto.Event
	CreateEventReqMock *proto.CreateEventRequest
	UpdateEventReqMock *proto.UpdateEventRequest
}

func TestEventService(t *testing.T) {
	suite.Run(t, new(EventServiceTest))
}

func (t *EventServiceTest) SetupTest() {
	t.Event = &event.Event{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		NameTH:        faker.Word(),
		DescriptionTH: faker.Paragraph(),
		NameEN:        faker.Word(),
		DescriptionEN: faker.Paragraph(),
		Code:          faker.Word(),
		ImageURL:      faker.Paragraph(),
	}

	t.EventDto = &proto.Event{
		Id:            t.Event.ID.String(),
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
		ImageURL:      t.Event.ImageURL,
	}

	t.CreateEventReqMock = &proto.CreateEventRequest{
		Event: &proto.Event{
			NameTH:        t.Event.NameTH,
			DescriptionTH: t.Event.DescriptionTH,
			NameEN:        t.Event.NameEN,
			DescriptionEN: t.Event.DescriptionEN,
			Code:          t.Event.Code,
			ImageURL:      t.Event.ImageURL,
		},
	}

	t.UpdateEventReqMock = &proto.UpdateEventRequest{
		Event: &proto.Event{
			Id:            t.Event.ID.String(),
			NameTH:        t.Event.NameTH,
			DescriptionTH: t.Event.DescriptionTH,
			NameEN:        t.Event.NameEN,
			DescriptionEN: t.Event.DescriptionEN,
			Code:          t.Event.Code,
			ImageURL:      t.Event.ImageURL,
		},
	}
}

func (t *EventServiceTest) createEventDto(in []*event.Event) []*proto.Event {
	var result []*proto.Event

	for _, e := range in {
		r := &proto.Event{
			Id:            e.ID.String(),
			NameTH:        e.NameTH,
			DescriptionTH: e.DescriptionTH,
			NameEN:        e.NameEN,
			DescriptionEN: e.DescriptionEN,
			Code:          e.Code,
			ImageURL:      e.ImageURL,
		}

		result = append(result, r)
	}

	return result
}

func (t *EventServiceTest) createEvent() []*event.Event {
	var result []*event.Event

	for i := 0; i <= 5; i++ {
		r := &event.Event{
			Base: model.Base{
				ID: uuid.New(),
			},
			NameTH:        faker.Word(),
			DescriptionTH: faker.Paragraph(),
			NameEN:        faker.Word(),
			DescriptionEN: faker.Paragraph(),
			Code:          faker.Word(),
			ImageURL:      faker.Paragraph(),
		}

		result = append(result, r)
	}

	return result
}

func (t *EventServiceTest) TestFindAllEventSuccess() {

	eventList := t.createEvent()

	want := &proto.FindAllEventResponse{Event: t.createEventDto(eventList)}

	var eventsIn []*event.Event

	r := mock.RepositoryMock{}
	r.On("FindAllEvent", &eventsIn).Return(eventList, nil)

	srv := NewService(&r)
	actual, err := srv.FindAllEvent(context.Background(), &proto.FindAllEventRequest{})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestFindEventByIDSuccess() {

	want := &proto.FindEventByIDResponse{Event: t.EventDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindEventByID", t.Event.ID.String(), &event.Event{}).Return(t.Event, nil)

	srv := NewService(repo)

	actual, err := srv.FindEventByID(context.Background(), &proto.FindEventByIDRequest{Id: t.Event.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestFindEventByIDNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindEventByID", t.Event.ID.String(), &event.Event{}).Return(nil, errors.New("Not found event"))

	srv := NewService(repo)

	actual, err := srv.FindEventByID(context.Background(), &proto.FindEventByIDRequest{Id: t.Event.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *EventServiceTest) TestCreateSuccess() {
	want := &proto.CreateEventResponse{Event: t.EventDto}

	repo := &mock.RepositoryMock{}

	in := &event.Event{
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
		ImageURL:      t.Event.ImageURL,
	}

	repo.On("Create", in).Return(t.Event, nil)
	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateEventReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestCreateInternalErr() {
	repo := &mock.RepositoryMock{}

	in := &event.Event{
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
		ImageURL:      t.Event.ImageURL,
	}

	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateEventReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *EventServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateEventResponse{Event: t.EventDto}

	eventIn := *t.Event

	repo := &mock.RepositoryMock{}

	repo.On("Update", t.Event.ID.String(), &eventIn).Return(t.Event, nil)

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateEventReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestUpdateNotFound() {
	eventIn := *t.Event

	repo := &mock.RepositoryMock{}
	repo.On("Update", t.Event.ID.String(), &eventIn).Return(nil, errors.New("Not found event"))

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateEventReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *EventServiceTest) TestUpdateMalformed() {
	eventIn := *t.Event

	repo := &mock.RepositoryMock{}
	repo.On("Update", t.Event.ID.String(), eventIn).Return(nil, errors.New("Not found event"))

	srv := NewService(repo)

	t.UpdateEventReqMock.Event.Id = "abc"

	actual, err := srv.Update(context.Background(), t.UpdateEventReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *EventServiceTest) TestDeleteSuccess() {
	want := &proto.DeleteEventResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.Event.ID.String()).Return(nil)

	srv := NewService(repo)
	actual, err := srv.Delete(context.Background(), &proto.DeleteEventRequest{Id: t.EventDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EventServiceTest) TestDeleteNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.Event.ID.String()).Return(errors.New("Not found event"))

	srv := NewService(repo)
	actual, err := srv.Delete(context.Background(), &proto.DeleteEventRequest{Id: t.EventDto.Id})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *EventServiceTest) TestFindAllEventWithTypeSuccess() {

	eventList := t.createEvent()

	want := &proto.FindAllEventWithTypeResponse{Event: t.createEventDto(eventList)}

	var eventsIn []*event.Event

	r := mock.RepositoryMock{}
	r.On("FindAllEventWithType", &eventsIn).Return(eventList, nil)

	srv := NewService(&r)
	actual, err := srv.FindAllEventWithType(context.Background(), &proto.FindAllEventWithTypeRequest{})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}
