package checkin

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Checkin(ci *checkin.Checkin) error {
	args := r.Called(ci)

	return args.Error(0)
}

func (r *RepositoryMock) Checkout(ci *checkin.Checkin) error {
	args := r.Called(ci)

	return args.Error(0)
}

func (r *RepositoryMock) FindLastCheckin(userid string, eventType int32, ci *checkin.Checkin) error {
	args := r.Called(userid, eventType, ci)

	if args.Get(0) != nil {
		*ci = *args.Get(0).(*checkin.Checkin)
	}

	return args.Error(1)
}
