package user

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) ConfirmEstamp(uId string, thisUser *user.User, thisEvent *event.Event) error {
	args := r.Called(uId, thisUser, thisEvent)
	if args.Get(0) != nil {
		*thisEvent = *args.Get(0).(*event.Event)
	}

	return args.Error(1)
}

func (r *RepositoryMock) GetUserEstamp(uId string, user *user.User, results *[]*event.Event) error {
	args := r.Called(uId, user, results)

	if args.Get(0) != nil {
		*results = args.Get(0).([]*event.Event)
	}

	return args.Error(1)
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

func (r *RepositoryMock) Verify(studentId string, verifyType string) error {
	args := r.Called(studentId, verifyType)

	return args.Error(0)
}

func (r *RepositoryMock) Update(id string, result *user.User) error {
	args := r.Called(id, result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) SaveUser(result *user.User) error {
	args := r.Called(result)

	if args.Get(0) != nil {
		*result = *args.Get(0).(*user.User)
	}

	return args.Error(1)
}

func (r *RepositoryMock) Delete(id string) error {
	args := r.Called(id)

	return args.Error(0)
}
