package user

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) CreateOrUpdate(in *user.User) error {
	args := r.Called(in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindOne(id string, result *user.User) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindByStudentID(sid string, result *user.User) error {
	args := r.Called(sid, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Create(in *user.User) error {
	args := r.Called(in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Verify(studentId string) error {
	args := r.Called(studentId)

	return args.Error(0)
}

func (r *RepositoryMock) Update(id string, result *user.User) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Delete(id string) error {
	args := r.Called(id)

	return args.Error(0)
}
