package group

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) RemoveAllBaan(in *group.Group) error {
	args := r.Called(in)

	return args.Error(0)
}

func (r *RepositoryMock) FindGroupWithBaans(token string, result *group.Group) error {
	args := r.Called(token, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*group.Group)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindGroupByToken(token string, result *group.Group) error {
	args := r.Called(token, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*group.Group)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindGroupById(id string, result *group.Group) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*group.Group)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Create(in *group.Group) error {
	args := r.Called(in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*group.Group)
	}

	return args.Error(1)
}

func (r *RepositoryMock) UpdateWithLeader(leaderId string, result *group.Group) error {
	args := r.Called(leaderId, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*group.Group)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Delete(id string) error {
	args := r.Called(id)

	return args.Error(0)
}
