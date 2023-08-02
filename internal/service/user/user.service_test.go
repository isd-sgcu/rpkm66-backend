package user

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/event"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/user"
	event_proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/event/v1"
	proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/user/v1"
	"github.com/isd-sgcu/rpkm66-backend/internal/utils"
	eMock "github.com/isd-sgcu/rpkm66-backend/mocks/event"
	fMock "github.com/isd-sgcu/rpkm66-backend/mocks/file"
	mock "github.com/isd-sgcu/rpkm66-backend/mocks/user"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServiceTest struct {
	suite.Suite
	User                         *user.User
	UpdateUser                   *user.User
	UserDto                      *proto.User
	CreateUserReqMock            *proto.CreateUserRequest
	UpdateUserReqMock            *proto.UpdateUserRequest
	UpdatePersonalityGameReqMock *proto.UpdatePersonalityGameRequest
	Event                        *event.Event
	EventDto                     *event_proto.Event
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.User = &user.User{
		Base: entity.Base{
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
		EmerPhone:       faker.Phonenumber(),
		EmerRelation:    faker.Word(),
		WantBottle:      getBoolPtr(),
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
		IsGotTicket:     utils.BoolAdr(true),
		GroupID:         utils.UUIDAdr(uuid.New()),
		BaanID:          utils.UUIDAdr(uuid.New()),
		PersonalityGame: faker.Word(),
	}

	t.Event = &event.Event{
		Base: entity.Base{
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

	t.EventDto = &event_proto.Event{
		Id:            t.Event.ID.String(),
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
		ImageURL:      t.Event.ImageURL,
	}

	t.Event = &event.Event{
		Base: entity.Base{
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

	t.EventDto = &event_proto.Event{
		Id:            t.Event.ID.String(),
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
		ImageURL:      t.Event.ImageURL,
	}

	t.Event = &event.Event{
		Base: entity.Base{
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

	t.EventDto = &event_proto.Event{
		Id:            t.Event.ID.String(),
		NameTH:        t.Event.NameTH,
		DescriptionTH: t.Event.DescriptionTH,
		NameEN:        t.Event.NameEN,
		DescriptionEN: t.Event.DescriptionEN,
		Code:          t.Event.Code,
		ImageURL:      t.Event.ImageURL,
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
		EmerPhone:       t.User.EmerPhone,
		EmerRelation:    t.User.EmerRelation,
		WantBottle:      *t.User.WantBottle,
		CanSelectBaan:   *t.User.CanSelectBaan,
		IsVerify:        *t.User.IsVerify,
		BaanId:          t.User.BaanID.String(),
		IsGotTicket:     *t.User.IsGotTicket,
		PersonalityGame: t.User.PersonalityGame,
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
			EmerPhone:       t.User.EmerPhone,
			EmerRelation:    t.User.EmerRelation,
			WantBottle:      *t.User.WantBottle,
			CanSelectBaan:   *t.User.CanSelectBaan,
			IsVerify:        *t.User.IsVerify,
			BaanId:          t.User.BaanID.String(),
			PersonalityGame: t.User.PersonalityGame,
		},
	}

	t.UpdateUserReqMock = &proto.UpdateUserRequest{
		Id:              t.User.ID.String(),
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		EmerPhone:       t.User.EmerPhone,
		EmerRelation:    t.User.EmerRelation,
		WantBottle:      *t.User.WantBottle,
		PersonalityGame: t.User.PersonalityGame,
	}

	t.UpdatePersonalityGameReqMock = &proto.UpdatePersonalityGameRequest{
		Id:              t.User.ID.String(),
		PersonalityGame: t.User.PersonalityGame,
	}

	t.UpdateUser = &user.User{
		Title:           t.User.Title,
		Firstname:       t.User.Firstname,
		Lastname:        t.User.Lastname,
		Nickname:        t.User.Nickname,
		Phone:           t.User.Phone,
		LineID:          t.User.LineID,
		Email:           t.User.Email,
		AllergyFood:     t.User.AllergyFood,
		FoodRestriction: t.User.FoodRestriction,
		AllergyMedicine: t.User.AllergyMedicine,
		Disease:         t.User.Disease,
		EmerPhone:       t.User.EmerPhone,
		EmerRelation:    t.User.EmerRelation,
		WantBottle:      t.User.WantBottle,
		PersonalityGame: t.User.PersonalityGame,
	}
}

func (t *UserServiceTest) TestConfirmEstampSuccess() {
	want := &proto.ConfirmEstampResponse{}

	repo := &mock.RepositoryMock{}
	fileSrv := &fMock.ServiceMock{}
	eventSrv := &eMock.RepositoryMock{}

	eventSrv.On("FindEventByID", t.Event.ID.String(), &event.Event{}).Return(t.Event, nil)
	repo.On("ConfirmEstamp", t.User.ID.String(), &user.User{
		Base: entity.Base{
			ID: t.User.ID,
		},
	}, t.Event).Return(nil, nil)

	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.ConfirmEstamp(context.Background(), &proto.ConfirmEstampRequest{
		UId: t.User.ID.String(),
		EId: t.Event.ID.String(),
	})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestGetUserEstampSuccess() {
	eventList := t.createEvent()
	want := &proto.GetUserEstampResponse{EventList: t.createEventDto(eventList)}

	var eventsIn []*event.Event

	r := mock.RepositoryMock{}
	r.On("GetUserEstamp", t.User.ID.String(), &user.User{
		Base: entity.Base{
			ID: t.User.ID,
		},
	}, &eventsIn).Return(eventList, nil)

	fileSrv := &fMock.ServiceMock{}
	eventSrv := &eMock.RepositoryMock{}

	srv := NewService(&r, fileSrv, eventSrv)
	actual, err := srv.GetUserEstamp(context.Background(), &proto.GetUserEstampRequest{
		UId: t.User.ID.String(),
	})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneSuccess() {
	url := faker.URL()

	t.UserDto.ImageUrl = url

	want := &proto.FindOneUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return(url, nil)

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneSignUrlErr() {
	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return("", status.Error(codes.Unavailable, "Cannot get signed url"))

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Unavailable, st.Code())
}

func (t *UserServiceTest) TestFindOneSignUrlNotFound() {
	want := &proto.FindOneUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.User.ID.String()).Return("", status.Error(codes.NotFound, "Not found file"))

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)

	actual, err := srv.FindOne(context.Background(), &proto.FindOneUserRequest{Id: t.User.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindOne", t.User.ID.String(), &user.User{}).Return(nil, errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)

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
		EmerPhone:       t.User.EmerPhone,
		EmerRelation:    t.User.EmerRelation,
		WantBottle:      t.User.WantBottle,
		CanSelectBaan:   t.User.CanSelectBaan,
		BaanID:          t.User.BaanID,
		PersonalityGame: t.User.PersonalityGame,
	}

	repo.On("Create", in).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)

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
		EmerPhone:       t.User.EmerPhone,
		EmerRelation:    t.User.EmerRelation,
		WantBottle:      t.User.WantBottle,
		CanSelectBaan:   t.User.CanSelectBaan,
		BaanID:          t.User.BaanID,
		PersonalityGame: t.User.PersonalityGame,
	}

	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)

	actual, err := srv.Create(context.Background(), t.CreateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *UserServiceTest) TestVerifySuccess() {
	want := &proto.VerifyUserResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Verify", t.User.ID.String(), "is_verify").Return(nil)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Verify(context.Background(), &proto.VerifyUserRequest{StudentId: t.UserDto.Id, VerifyType: "vaccine"})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestVerifyInvalidArgument() {
	repo := &mock.RepositoryMock{}

	repo.On("Verify", t.User.ID.String(), "is_verify").Return(gorm.ErrRecordNotFound)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Verify(context.Background(), &proto.VerifyUserRequest{StudentId: t.UserDto.Id, VerifyType: "-"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *UserServiceTest) TestVerifyNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Verify", t.User.ID.String(), "is_verify").Return(gorm.ErrRecordNotFound)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Verify(context.Background(), &proto.VerifyUserRequest{StudentId: t.UserDto.Id, VerifyType: "vaccine"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("Update", t.User.ID.String(), t.UpdateUser).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestUpdateNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.User.ID.String(), t.UpdateUser).Return(nil, errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Update(context.Background(), t.UpdateUserReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestDeleteSuccess() {
	want := &proto.DeleteUserResponse{Success: true}

	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.User.ID.String()).Return(nil)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Delete(context.Background(), &proto.DeleteUserRequest{Id: t.UserDto.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestDeleteNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("Delete", t.User.ID.String()).Return(errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.Delete(context.Background(), &proto.DeleteUserRequest{Id: t.UserDto.Id})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) TestCreateOrUpdateSuccess() {
	userIn := *t.User
	userIn.IsVerify = nil
	userIn.GroupID = nil
	userIn.IsGotTicket = nil
	want := &proto.CreateOrUpdateUserResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", &userIn).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestCreateOrUpdateMalformedID() {
	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", t.User).Return(t.User, nil)

	t.UserDto.Id = "abc"

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.CreateOrUpdate(context.Background(), &proto.CreateOrUpdateUserRequest{User: t.UserDto})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *UserServiceTest) TestCreateOrUpdateInternalErr() {
	userIn := *t.User
	userIn.IsVerify = nil
	userIn.GroupID = nil
	userIn.IsGotTicket = nil
	repo := &mock.RepositoryMock{}

	repo.On("CreateOrUpdate", &userIn).Return(nil, errors.New("Something wrong"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
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

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.FindByStudentID(context.Background(), &proto.FindByStudentIDUserRequest{StudentId: t.UserDto.StudentID})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindByStudentIDNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindByStudentID", t.User.StudentID, &user.User{}).Return(nil, errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.FindByStudentID(context.Background(), &proto.FindByStudentIDUserRequest{StudentId: t.UserDto.StudentID})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *UserServiceTest) createEventDto(in []*event.Event) []*event_proto.Event {
	var result []*event_proto.Event

	for _, e := range in {
		r := &event_proto.Event{
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

func (t *UserServiceTest) createEvent() []*event.Event {
	var result []*event.Event

	for i := 0; i <= 5; i++ {
		r := &event.Event{
			Base: entity.Base{
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

func getBoolPtr() *bool {
	value := rand.Intn(2) == 0
	return &value
}
func (t *UserServiceTest) TestUpdatePersonalityGameSuccess() {
	want := &proto.UpdatePersonalityGameResponse{User: t.UserDto}

	repo := &mock.RepositoryMock{}
	repo.On("Update", t.User.ID.String(), &user.User{
		PersonalityGame: t.UpdateUser.PersonalityGame,
	}).Return(t.User, nil)

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.UpdatePersonalityGame(context.Background(), t.UpdatePersonalityGameReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestUpdatePersonalityGameNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("Update", t.User.ID.String(), &user.User{
		PersonalityGame: t.UpdateUser.PersonalityGame,
	}).Return(nil, errors.New("Not found user"))

	fileSrv := &fMock.ServiceMock{}

	eventSrv := &eMock.RepositoryMock{}
	srv := NewService(repo, fileSrv, eventSrv)
	actual, err := srv.UpdatePersonalityGame(context.Background(), t.UpdatePersonalityGameReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}
