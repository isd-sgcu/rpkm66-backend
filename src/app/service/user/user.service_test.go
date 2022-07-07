package user

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/user"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"testing"
	"time"
)

type UserServiceTest struct {
	suite.Suite
	User              *user.User
	UserDto           *proto.User
	CreateUserReqMock *proto.CreateUserRequest
	UpdateUserReqMock *proto.UpdateUserRequest
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.User = &user.User{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Title:           faker.Word(),
		Firstname:       faker.FirstName(),
		Lastname:        faker.LastName(),
		Nickname:        faker.Name(),
		StudentID:       faker.Word(),
		Faculty:         faker.Word(),
		Year:            faker.Word(),
		Phone:           faker.Phonenumber(),
		LineID:          faker.Word(),
		Email:           faker.Email(),
		AllergyFood:     faker.Word(),
		FoodRestriction: faker.Word(),
		AllergyMedicine: faker.Word(),
		Disease:         faker.Word(),
		ImageUrl:        faker.URL(),
		CanSelectBaan:   true,
	}

	t.UserDto = &proto.User{
		Id:              t.User.ID.String(),
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		ImageUrl:        t.User.ImageUrl,
		CanSelectBaan:   t.User.CanSelectBaan,
	}

	t.CreateUserReqMock = &proto.CreateUserRequest{
		User: &proto.User{
			Title:           t.User.Title,
			Firstname:       t.User.Firstname,
			Lastname:        t.User.Lastname,
			Nickname:        t.User.Nickname,
			StudentID:       t.User.StudentID,
			Faculty:         t.User.Faculty,
			Year:            t.User.Year,
			Phone:           t.User.Phone,
			LineID:          t.User.LineID,
			Email:           t.User.Email,
			AllergyFood:     t.User.AllergyFood,
			FoodRestriction: t.User.FoodRestriction,
			AllergyMedicine: t.User.AllergyMedicine,
			Disease:         t.User.Disease,
			ImageUrl:        t.User.ImageUrl,
			CanSelectBaan:   t.User.CanSelectBaan,
		},
	}

	t.UpdateUserReqMock = &proto.UpdateUserRequest{
		User: &proto.User{
			Id:              t.User.ID.String(),
			Title:           t.User.Title,
			Firstname:       t.User.Firstname,
			Lastname:        t.User.Lastname,
			Nickname:        t.User.Nickname,
			StudentID:       t.User.StudentID,
			Faculty:         t.User.Faculty,
			Year:            t.User.Year,
			Phone:           t.User.Phone,
			LineID:          t.User.LineID,
			Email:           t.User.Email,
			AllergyFood:     t.User.AllergyFood,
			FoodRestriction: t.User.FoodRestriction,
			AllergyMedicine: t.User.AllergyMedicine,
			Disease:         t.User.Disease,
			ImageUrl:        t.User.ImageUrl,
			CanSelectBaan:   t.User.CanSelectBaan,
		},
	}
}

func (t *UserServiceTest) TestFindOneSuccess() {
	want := &proto.FindOneUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	srv := NewService(repo)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(nil, errors.New("Not found user"))

	srv := NewService(repo)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestCreateSuccess() {
	want := &proto.CreateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	in := &user.User{
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		ImageUrl:        t.User.ImageUrl,
		CanSelectBaan:   t.User.CanSelectBaan,
	}

	repo.On("Create", in).Return(t.User, nil)

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateUserReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestCreateInternalErr() {
	repo := &mock.RepositoryMock{}

	in := &user.User{
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		StudentID:       t.User.StudentID,
		Faculty:         t.User.Faculty,
		Year:            t.User.Year,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		ImageUrl:        t.User.ImageUrl,
		CanSelectBaan:   t.User.CanSelectBaan,
	}

	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *UserServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("Update", t.User.ID.String(), t.User).Return(t.User, nil)

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestUpdateNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.User.ID.String(), t.User).Return(nil, errors.New("Not found user"))

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestUpdateMalformed() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.User.ID.String(), t.User).Return(nil, errors.New("Not found user"))

	srv := NewService(repo)

	t.UpdateUserReqMock.User.Id = "abc"

	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *UserServiceTest) TestDeleteSuccess() {
	want := &proto.DeleteUserResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.User.ID.String()).Return(nil)

	srv := NewService(repo)
	actual, err := srv.Delete(context.Background(), &proto.DeleteUserRequest{Id: t.UserDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestDeleteNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.User.ID.String()).Return(errors.New("Not found user"))

	srv := NewService(repo)
	actual, err := srv.Delete(context.Background(), &proto.DeleteUserRequest{Id: t.UserDto.Id})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestCreateOrUpdateSuccess() {
	want := &proto.CreateOrUpdateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", t.User).Return(t.User, nil)

	srv := NewService(repo)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestCreateOrUpdateMalformedID() {
	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", t.User).Return(t.User, nil)

	t.UserDto.Id = "abc"

	srv := NewService(repo)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *UserServiceTest) TestCreateOrUpdateInternalErr() {
	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", t.User).Return(nil, errors.New("Something wrong"))

	srv := NewService(repo)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *UserServiceTest) TestFindByStudentIDSuccess() {
	want := &proto.FindByStudentIDUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("FindByStudentID", t.User.StudentID, &user.User{}).Return(t.User, nil)

	srv := NewService(repo)
	actual, err := srv.FindByStudentID(context.Background(), &proto.FindByStudentIDUserRequest{StudentId: t.UserDto.StudentID})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindByStudentIDNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindByStudentID", t.User.StudentID, &user.User{}).Return(nil, errors.New("Not found user"))

	srv := NewService(repo)
	actual, err := srv.FindByStudentID(context.Background(), &proto.FindByStudentIDUserRequest{StudentId: t.UserDto.StudentID})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}
