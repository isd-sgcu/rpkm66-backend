package event

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) FindAllEvent(result *[]*event.Event) error {
	args := r.Called(result)

	if args.Get(0) != nil {
		*result = args.Get(0).([]*event.Event)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindEventByID(id string, result *event.Event) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*event.Event)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Create(in *event.Event) error {
	args := r.Called(in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*event.Event)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Update(id string, result *event.Event) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*event.Event)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Delete(id string) error {
	args := r.Called(id)

	return args.Error(0)
}

func (r *RepositoryMock) FindAllEventWithType(eventType string, result *[]*event.Event) error {
	args := r.Called(result)

	if args.Get(0) != nil {
		*result = args.Get(0).([]*event.Event)
	}

	return args.Error(1)
}
