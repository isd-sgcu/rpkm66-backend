package group

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) FindUserById(id string, result *user.User) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
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

func (r *RepositoryMock) Create(in *group.Group) error {
	args := r.Called(in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*group.Group)
	}

	return args.Error(1)
}

func (r *RepositoryMock) UpdateUser(result *user.User) error {
	args := r.Called(result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
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
