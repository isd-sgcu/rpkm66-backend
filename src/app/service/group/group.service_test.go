package group

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/baan"
	baan_group_selection "github.com/isd-sgcu/rpkm66-backend/src/app/model/baan-group-selection"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/group"
	"github.com/isd-sgcu/rpkm66-backend/src/app/model/user"
	"github.com/isd-sgcu/rpkm66-backend/src/app/utils"
	"github.com/isd-sgcu/rpkm66-backend/src/config"
	size "github.com/isd-sgcu/rpkm66-backend/src/constant/baan"
	mockBaan "github.com/isd-sgcu/rpkm66-backend/src/mocks/baan"
	mockBgs "github.com/isd-sgcu/rpkm66-backend/src/mocks/baan-group-selection"
	mockCache "github.com/isd-sgcu/rpkm66-backend/src/mocks/cache"
	mockFile "github.com/isd-sgcu/rpkm66-backend/src/mocks/file"
	mock "github.com/isd-sgcu/rpkm66-backend/src/mocks/group"
	mockUser "github.com/isd-sgcu/rpkm66-backend/src/mocks/user"
	"github.com/isd-sgcu/rpkm66-backend/src/proto"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"testing"
	"time"
)

type GroupServiceTest struct {
	suite.Suite
	Group              *group.Group
	GroupDto           *proto.Group
	Member             *user.User
	UpdateGroupReqMock *proto.UpdateGroupRequest
	UserMock           *user.User
	UserDtoMock        *proto.UserInfo
	ReservedUser       *user.User
	RemovedUser        *user.User
	PreviousGroup      *group.Group
	conf               config.App
}

func TestGroupService(t *testing.T) {
	suite.Run(t, new(GroupServiceTest))
}

func (t *GroupServiceTest) SetupTest() {
	t.UserMock = &user.User{
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
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
		GroupID:         utils.UUIDAdr(uuid.New()),
	}
	t.UserDtoMock = &proto.UserInfo{
		Id:        t.UserMock.ID.String(),
		Firstname: t.UserMock.Firstname,
		Lastname:  t.UserMock.Lastname,
		ImageUrl:  "",
	}
	t.Group = &group.Group{
		Base: model.Base{
			ID:        *t.UserMock.GroupID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: t.UserMock.ID.String(),
		Token:    faker.Word(),
		Members:  []*user.User{t.UserMock},
	}

	t.Member = &user.User{
		Base: model.Base{
			ID:        t.UserMock.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Firstname: t.UserMock.Firstname,
		Lastname:  t.UserMock.Lastname,
	}

	t.GroupDto = &proto.Group{
		Id:       t.Group.ID.String(),
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*proto.UserInfo{t.UserDtoMock},
	}

	t.UpdateGroupReqMock = &proto.UpdateGroupRequest{
		Group: &proto.Group{
			Id:       t.Group.ID.String(),
			LeaderID: t.Group.LeaderID,
			Token:    t.Group.Token,
			Members:  t.GroupDto.Members,
		},
		LeaderId: t.Group.LeaderID,
	}

	t.ReservedUser = &user.User{
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
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
		GroupID:         utils.UUIDAdr(uuid.New()),
	}

	t.RemovedUser = &user.User{
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
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}

	t.PreviousGroup = &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{t.UserMock, t.RemovedUser},
	}

	t.conf = config.App{
		Port:         3001,
		Debug:        false,
		BaanCacheTTL: 900,
		NBaan:        3,
	}
}

func createBaan() []*baan.Baan {
	var baans []*baan.Baan

	for i := 0; i < 3; i++ {
		b := baan.Baan{
			Base:          model.Base{ID: uuid.New()},
			NameTH:        faker.Word(),
			DescriptionTH: faker.Word(),
			NameEN:        faker.Word(),
			DescriptionEN: faker.Word(),
			ImageUrl:      faker.URL(),
			Size:          size.M,
			Facebook:      faker.Word(),
			FacebookUrl:   faker.URL(),
			Instagram:     faker.Word(),
			InstagramUrl:  faker.URL(),
			Line:          faker.Word(),
			LineUrl:       faker.URL(),
		}

		baans = append(baans, &b)
	}
	return baans
}

func createBaansDto() []*proto.Baan {
	var baans []*proto.Baan

	for i := 0; i < 3; i++ {
		b := proto.Baan{
			Id:            faker.UUIDDigit(),
			NameTH:        faker.Word(),
			DescriptionTH: faker.Paragraph(),
			NameEN:        faker.Word(),
			DescriptionEN: faker.Paragraph(),
			Size:          size.M,
			Facebook:      faker.URL(),
			Instagram:     faker.URL(),
			Line:          faker.URL(),
		}

		baans = append(baans, &b)
	}

	return baans
}

func (t *GroupServiceTest) TestFindOneSuccess() {
	baans := createBaan()
	baansSelection, _ := createBaansArray(t.Group.ID, baans)

	want := &proto.FindOneGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", (*t.UserMock.GroupID).String(), &group.Group{}).Return(t.Group, nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	var baansSelectionIn []*baan_group_selection.BaanGroupSelection
	bgsRepo := &mockBgs.RepositoryMock{}
	bgsRepo.On("FindBaans", t.Group.ID.String(), &baansSelectionIn).Return(baansSelection, nil)

	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.UserMock.ID.String()).Return("", nil)

	var baanIns []*proto.BaanInfo
	cacheRepo := &mockCache.RepositoryMock{}
	cacheRepo.On("GetCache", t.Group.ID.String(), &baanIns).Return(baans, nil)

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.FindOne(context.Background(), &proto.FindOneGroupRequest{UserId: t.UserMock.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindOneNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", (*t.UserMock.GroupID).String(), &group.Group{}).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.FindOne(context.Background(), &proto.FindOneGroupRequest{UserId: t.UserMock.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestFindOneInvalidID() {
	repo := &mock.RepositoryMock{}

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.FindOne(context.Background(), &proto.FindOneGroupRequest{UserId: "abc"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestFindOneWithCreateGroup() {
	want := &proto.FindOneGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	in := &group.Group{
		LeaderID: t.UserMock.ID.String(),
	}
	repo.On("Create", in).Return(t.Group, nil)
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)
	nonGroupUser := &user.User{
		Base: model.Base{
			ID:        t.UserMock.ID,
			CreatedAt: t.UserMock.CreatedAt,
			UpdatedAt: t.UserMock.UpdatedAt,
			DeletedAt: t.UserMock.DeletedAt,
		},
		Title:           t.UserMock.Title,
		Firstname:       t.UserMock.Firstname,
		Lastname:        t.UserMock.Lastname,
		Nickname:        t.UserMock.Nickname,
		StudentID:       t.UserMock.StudentID,
		Faculty:         t.UserMock.Faculty,
		Year:            t.UserMock.Year,
		Phone:           t.UserMock.Phone,
		LineID:          t.UserMock.LineID,
		Email:           t.UserMock.Email,
		AllergyFood:     t.UserMock.AllergyFood,
		FoodRestriction: t.UserMock.FoodRestriction,
		AllergyMedicine: t.UserMock.AllergyMedicine,
		Disease:         t.UserMock.Disease,
		CanSelectBaan:   t.UserMock.CanSelectBaan,
		IsVerify:        t.UserMock.IsVerify,
		GroupID:         nil,
	}

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(nonGroupUser, nil)
	userRepo.On("Update", t.UserMock.ID.String(), t.UserMock).Return(t.UserMock, nil)
	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.UserMock.ID.String()).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.FindOne(context.Background(), &proto.FindOneGroupRequest{UserId: t.UserMock.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindByTokenSuccess() {
	want := &proto.FindByTokenGroupResponse{
		Id:    t.Group.ID.String(),
		Token: t.Group.Token,
		Leader: &proto.UserInfo{
			Id:        t.Group.LeaderID,
			Firstname: t.UserMock.Firstname,
			Lastname:  t.UserMock.Lastname,
			ImageUrl:  "",
		},
	}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.UserMock.ID.String()).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.FindByToken(context.Background(), &proto.FindByTokenGroupRequest{Token: t.GroupDto.Token})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindByTokenNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.FindByToken(context.Background(), &proto.FindByTokenGroupRequest{Token: t.GroupDto.Token})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateGroupResponse{Group: t.GroupDto}

	nonUser := &user.User{
		Base: model.Base{
			ID:        t.UserMock.ID,
			CreatedAt: t.UserMock.CreatedAt,
			UpdatedAt: t.UserMock.UpdatedAt,
			DeletedAt: t.UserMock.DeletedAt,
		},
		Firstname: t.UserMock.Firstname,
		Lastname:  t.UserMock.Lastname,
	}

	raw := &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{nonUser},
	}

	repo := &mock.RepositoryMock{}
	repo.On("UpdateWithLeader", t.Group.LeaderID, raw).Return(raw, nil)

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.Group.LeaderID).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestUpdateNotFound() {
	t.Group.Members = []*user.User{t.Member}

	repo := &mock.RepositoryMock{}
	repo.On("UpdateWithLeader", t.Group.LeaderID, t.Group).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestUpdateMalformed() {
	repo := &mock.RepositoryMock{}
	repo.On("UpdateWithLeader", t.Group.LeaderID, t.Group).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	t.UpdateGroupReqMock.Group.Id = "abc"

	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

// Case1 : a user is a not a king in the group --> expected result : the user is able to join other group, and remove the user from the previous group
func (t *GroupServiceTest) TestJoinSuccess1() {
	afterJoinedUser := &user.User{
		Base: model.Base{
			ID:        t.ReservedUser.ID,
			CreatedAt: t.ReservedUser.CreatedAt,
			UpdatedAt: t.ReservedUser.UpdatedAt,
			DeletedAt: t.ReservedUser.DeletedAt,
		},
		Title:           t.ReservedUser.Title,
		Firstname:       t.ReservedUser.Firstname,
		Lastname:        t.ReservedUser.Lastname,
		Nickname:        t.ReservedUser.Nickname,
		StudentID:       t.ReservedUser.StudentID,
		Faculty:         t.ReservedUser.Faculty,
		Year:            t.ReservedUser.Year,
		Phone:           t.ReservedUser.Phone,
		LineID:          t.ReservedUser.LineID,
		Email:           t.ReservedUser.Email,
		AllergyFood:     t.ReservedUser.AllergyFood,
		FoodRestriction: t.ReservedUser.FoodRestriction,
		AllergyMedicine: t.ReservedUser.AllergyMedicine,
		Disease:         t.ReservedUser.Disease,
		CanSelectBaan:   t.ReservedUser.CanSelectBaan,
		IsVerify:        t.ReservedUser.IsVerify,
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}

	prevGrp := &group.Group{
		Base: model.Base{
			ID:        *t.ReservedUser.GroupID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: t.UserMock.ID.String(),
		Token:    faker.Word(),
		Members:  []*user.User{t.UserMock, t.ReservedUser},
	}

	joinGrp := &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{t.UserMock, t.UserMock},
	}

	want := &proto.JoinGroupResponse{Group: &proto.Group{
		Id:       t.Group.ID.String(),
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*proto.UserInfo{t.UserDtoMock, t.UserDtoMock},
	}}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(joinGrp, nil)
	repo.On("FindGroupById", (*t.ReservedUser.GroupID).String(), &group.Group{}).Return(prevGrp, nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.ReservedUser.ID.String(), &user.User{}).Return(t.ReservedUser, nil)
	userRepo.On("Update", t.ReservedUser.ID.String(), afterJoinedUser).Return(t.UserMock, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.UserMock.ID.String()).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.Group.Token, UserId: t.ReservedUser.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

// Case2 : a user is only one in the group --> expected result : the user is able to join other group, and delete the previous group
func (t *GroupServiceTest) TestJoinSuccess2() {
	headUserDto := &proto.UserInfo{
		Id:        t.ReservedUser.ID.String(),
		Firstname: t.ReservedUser.Firstname,
		Lastname:  t.ReservedUser.Lastname,
		ImageUrl:  "",
	}

	joinUser := &user.User{
		Base: model.Base{
			ID:        t.UserMock.ID,
			CreatedAt: t.UserMock.CreatedAt,
			UpdatedAt: t.UserMock.UpdatedAt,
			DeletedAt: t.UserMock.DeletedAt,
		},
		Title:           t.UserMock.Title,
		Firstname:       t.UserMock.Firstname,
		Lastname:        t.UserMock.Lastname,
		Nickname:        t.UserMock.Nickname,
		StudentID:       t.UserMock.StudentID,
		Faculty:         t.UserMock.Faculty,
		Year:            t.UserMock.Year,
		Phone:           t.UserMock.Phone,
		LineID:          t.UserMock.LineID,
		Email:           t.UserMock.Email,
		AllergyFood:     t.UserMock.AllergyFood,
		FoodRestriction: t.UserMock.FoodRestriction,
		AllergyMedicine: t.UserMock.AllergyMedicine,
		Disease:         t.UserMock.Disease,
		CanSelectBaan:   t.UserMock.CanSelectBaan,
		IsVerify:        t.UserMock.IsVerify,
		GroupID:         t.ReservedUser.GroupID,
	}

	joinGroup := &group.Group{
		Base: model.Base{
			ID:        *t.ReservedUser.GroupID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: t.ReservedUser.ID.String(),
		Token:    faker.Word(),
		Members:  []*user.User{t.ReservedUser, t.ReservedUser},
	}

	want := &proto.JoinGroupResponse{Group: &proto.Group{
		Id:       joinGroup.ID.String(),
		LeaderID: joinGroup.LeaderID,
		Token:    joinGroup.Token,
		Members:  []*proto.UserInfo{headUserDto, headUserDto},
	}}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", joinGroup.Token, &group.Group{}).Return(joinGroup, nil)
	repo.On("FindGroupById", (*t.UserMock.GroupID).String(), &group.Group{}).Return(t.Group, nil)
	repo.On("Delete", (*t.UserMock.GroupID).String()).Return(nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserDtoMock.Id, &user.User{}).Return(t.UserMock, nil)
	userRepo.On("Update", joinUser.ID.String(), joinUser).Return(joinUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.ReservedUser.ID.String()).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: joinGroup.Token, UserId: t.UserDtoMock.Id})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

// Case3 : a member of the group can not join their own group
func (t *GroupServiceTest) TestJoinForbidden() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.Group.Token, UserId: t.UserDtoMock.Id})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

// Wrong Token
func (t *GroupServiceTest) TestJoinNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.GroupDto.Token, UserId: uuid.New().String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

// Wrong userId
func (t *GroupServiceTest) TestJoinMalformed() {
	repo := &mock.RepositoryMock{}

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.GroupDto.Token, UserId: "abc"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestJoinFullGroup() {
	member1User := &user.User{
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
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
		GroupID:         t.ReservedUser.GroupID,
	}
	member2User := &user.User{
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
		CanSelectBaan:   utils.BoolAdr(true),
		IsVerify:        utils.BoolAdr(true),
		GroupID:         t.ReservedUser.GroupID,
	}

	fullGroup := &group.Group{
		Base: model.Base{
			ID:        *t.ReservedUser.GroupID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: t.ReservedUser.ID.String(),
		Token:    faker.Word(),
		Members:  []*user.User{t.ReservedUser, member1User, member2User},
	}
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", fullGroup.Token, &group.Group{}).Return(fullGroup, nil)

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: fullGroup.Token, UserId: t.UserDtoMock.Id})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *GroupServiceTest) TestDeleteMemberSuccess() {
	in := &group.Group{
		LeaderID: t.RemovedUser.ID.String(),
	}
	createdUser := &user.User{
		Base: model.Base{
			ID:        t.RemovedUser.ID,
			CreatedAt: t.RemovedUser.CreatedAt,
			UpdatedAt: t.RemovedUser.UpdatedAt,
			DeletedAt: t.RemovedUser.DeletedAt,
		},
		Title:           t.RemovedUser.Title,
		Firstname:       t.RemovedUser.Firstname,
		Lastname:        t.RemovedUser.Lastname,
		Nickname:        t.RemovedUser.Nickname,
		StudentID:       t.RemovedUser.StudentID,
		Faculty:         t.RemovedUser.Faculty,
		Year:            t.RemovedUser.Year,
		Phone:           t.RemovedUser.Phone,
		LineID:          t.RemovedUser.LineID,
		Email:           t.RemovedUser.Email,
		AllergyFood:     t.RemovedUser.AllergyFood,
		FoodRestriction: t.RemovedUser.FoodRestriction,
		AllergyMedicine: t.RemovedUser.AllergyMedicine,
		Disease:         t.RemovedUser.Disease,
		CanSelectBaan:   t.ReservedUser.CanSelectBaan,
		IsVerify:        t.RemovedUser.IsVerify,
		GroupID:         utils.UUIDAdr(in.ID),
	}
	want := &proto.DeleteMemberGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}
	repo.On("Create", in).Return(in, nil)
	repo.On("FindGroupById", t.RemovedUser.GroupID.String(), &group.Group{}).Return(t.PreviousGroup, nil)
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.RemovedUser.ID.String(), &user.User{}).Return(t.RemovedUser, nil)
	userRepo.On("Update", createdUser.ID.String(), createdUser).Return(createdUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.UserMock.ID.String()).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{UserId: t.RemovedUser.ID.String(), LeaderId: t.GroupDto.LeaderID})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestDeleteMemberForbidden() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", t.RemovedUser.GroupID.String(), &group.Group{}).Return(t.PreviousGroup, nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.RemovedUser.ID.String(), &user.User{}).Return(t.RemovedUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{UserId: t.RemovedUser.ID.String(), LeaderId: t.RemovedUser.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *GroupServiceTest) TestDeleteMemberNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", t.RemovedUser.GroupID.String(), &group.Group{}).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.RemovedUser.ID.String(), &user.User{}).Return(t.RemovedUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)
	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{UserId: t.RemovedUser.ID.String(), LeaderId: t.GroupDto.LeaderID})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestDeleteMemberMalformed() {
	repo := &mock.RepositoryMock{}
	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{UserId: "abc", LeaderId: t.GroupDto.LeaderID})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupSuccess() {
	repo := &mock.RepositoryMock{}

	in := &group.Group{
		LeaderID: t.RemovedUser.ID.String(),
	}

	repo.On("Create", in).Return(in, nil)

	updatedUser := &user.User{
		Base: model.Base{
			ID:        t.RemovedUser.ID,
			CreatedAt: t.RemovedUser.CreatedAt,
			UpdatedAt: t.RemovedUser.UpdatedAt,
			DeletedAt: t.RemovedUser.DeletedAt,
		},
		Title:           t.RemovedUser.Title,
		Firstname:       t.RemovedUser.Firstname,
		Lastname:        t.RemovedUser.Lastname,
		Nickname:        t.RemovedUser.Nickname,
		StudentID:       t.RemovedUser.StudentID,
		Faculty:         t.RemovedUser.Faculty,
		Year:            t.RemovedUser.Year,
		Phone:           t.RemovedUser.Phone,
		LineID:          t.RemovedUser.LineID,
		Email:           t.RemovedUser.Email,
		AllergyFood:     t.RemovedUser.AllergyFood,
		FoodRestriction: t.RemovedUser.FoodRestriction,
		AllergyMedicine: t.RemovedUser.AllergyMedicine,
		Disease:         t.RemovedUser.Disease,
		CanSelectBaan:   t.RemovedUser.CanSelectBaan,
		IsVerify:        t.RemovedUser.IsVerify,
		GroupID:         utils.UUIDAdr(in.ID),
	}
	updatedUserDto := &proto.UserInfo{
		Id:        updatedUser.ID.String(),
		Firstname: updatedUser.Firstname,
		Lastname:  updatedUser.Lastname,
		ImageUrl:  "",
	}

	newGroup := &group.Group{
		Base: model.Base{
			ID:        in.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: in.LeaderID,
		Token:    in.Token,
		Members:  []*user.User{updatedUser},
	}

	repo.On("FindGroupById", t.RemovedUser.GroupID.String(), &group.Group{}).Return(t.PreviousGroup, nil)
	repo.On("FindGroupByToken", newGroup.Token, &group.Group{}).Return(newGroup, nil)

	want := &proto.LeaveGroupResponse{
		Group: &proto.Group{
			Id:       in.ID.String(),
			LeaderID: in.LeaderID,
			Token:    in.Token,
			Members:  []*proto.UserInfo{updatedUserDto},
		},
	}

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.RemovedUser.ID.String(), &user.User{}).Return(t.RemovedUser, nil)
	userRepo.On("Update", updatedUser.ID.String(), updatedUser).Return(updatedUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}
	fileSrv.On("GetSignedUrl", t.RemovedUser.ID.String()).Return("", nil)

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{UserId: t.RemovedUser.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestLeaveGroupNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", t.RemovedUser.GroupID.String(), &group.Group{}).Return(nil, errors.New("not found group"))

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.RemovedUser.ID.String(), &user.User{}).Return(t.RemovedUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{UserId: t.RemovedUser.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupMalformed() {
	repo := &mock.RepositoryMock{}

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{UserId: "abc"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupForbidden() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", t.UserMock.GroupID.String(), &group.Group{}).Return(t.PreviousGroup, nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{UserId: t.UserMock.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupInternalErr() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupById", t.RemovedUser.GroupID.String(), &group.Group{}).Return(t.PreviousGroup, nil)

	in := &group.Group{
		LeaderID: t.RemovedUser.ID.String(),
	}
	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.RemovedUser.ID.String(), &user.User{}).Return(t.RemovedUser, nil)

	bgsRepo := &mockBgs.RepositoryMock{}
	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{UserId: t.RemovedUser.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func createBaansArray(groupId uuid.UUID, baan []*baan.Baan) ([]*baan_group_selection.BaanGroupSelection, []string) {
	var baanSelections []*baan_group_selection.BaanGroupSelection
	var baanIds []string
	baans := createBaansDto()
	for order, b := range baans {
		baanSelection := &baan_group_selection.BaanGroupSelection{
			BaanID:  utils.UUIDAdr(uuid.MustParse(b.Id)),
			GroupID: utils.UUIDAdr(groupId),
			Order:   order + 1,
			Baan:    baan[order],
		}

		baanIds = append(baanIds, b.Id)
		baanSelections = append(baanSelections, baanSelection)
	}

	return baanSelections, baanIds
}

func createBaanInfo(baans []*baan.Baan) []*proto.BaanInfo {
	var baanInfos []*proto.BaanInfo
	for _, baan := range baans {

		b := &proto.BaanInfo{
			Id:       baan.ID.String(),
			NameTH:   baan.NameTH,
			NameEN:   baan.NameEN,
			ImageUrl: baan.ImageUrl,
		}

		baanInfos = append(baanInfos, b)
	}

	return baanInfos
}

func (t *GroupServiceTest) TestUpdateBaanSelectionSuccess() {
	want := &proto.SelectBaanResponse{Success: true}

	baans := createBaan()
	baansSelection, baanIds := createBaansArray(t.Group.ID, baans)
	baanInfos := createBaanInfo(baans)

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupWithBaans", t.GroupDto.Id, &group.Group{}).Return(t.Group, nil)
	repo.On("RemoveAllBaan", &group.Group{Base: model.Base{ID: t.Group.ID}}).Return(nil)

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	var baansSelectionIn []*baan_group_selection.BaanGroupSelection
	baanSelectionSaveIn := []*baan_group_selection.BaanGroupSelection{
		{
			BaanID:  baansSelection[0].BaanID,
			GroupID: baansSelection[0].GroupID,
			Order:   1,
		},
		{
			BaanID:  baansSelection[1].BaanID,
			GroupID: baansSelection[1].GroupID,
			Order:   2,
		},
		{
			BaanID:  baansSelection[2].BaanID,
			GroupID: baansSelection[2].GroupID,
			Order:   3,
		},
	}
	bgsRepo := &mockBgs.RepositoryMock{}
	bgsRepo.On("SaveBaansSelection", &baanSelectionSaveIn).Return(baansSelection, nil)
	bgsRepo.On("FindBaans", t.Group.ID.String(), &baansSelectionIn).Return(baansSelection, nil)

	cacheRepo := &mockCache.RepositoryMock{
		V: map[string]interface{}{},
	}
	cacheRepo.On("SaveCache", t.Group.ID.String(), &baanInfos, t.conf.BaanCacheTTL).Return(nil)

	fileSrv := &mockFile.ServiceMock{}

	var baanResult []*baan.Baan
	baanRepo := &mockBaan.RepositoryMock{}
	baanRepo.On("FindMulti", baanIds, &baanResult).Return(baans, nil)

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.SelectBaan(context.Background(), &proto.SelectBaanRequest{UserId: t.UserMock.ID.String(), Baans: baanIds})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestUpdateBaanSelectionNotFoundGroup() {
	baans := createBaan()
	_, baanIds := createBaansArray(t.Group.ID, baans)

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupWithBaans", t.GroupDto.Id, &group.Group{}).Return(nil, errors.New("Not found group"))

	userRepo := &mockUser.RepositoryMock{}
	userRepo.On("FindOne", t.UserMock.ID.String(), &user.User{}).Return(t.UserMock, nil)

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.SelectBaan(context.Background(), &proto.SelectBaanRequest{UserId: t.UserMock.ID.String(), Baans: baanIds})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestUpdateBaanSelectionDuplicatedBaan() {
	baans := createBaan()
	_, baanIds := createBaansArray(t.Group.ID, baans)
	baanIds[1] = baanIds[2]

	repo := &mock.RepositoryMock{}

	userRepo := &mockUser.RepositoryMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, t.conf)

	actual, err := srv.SelectBaan(context.Background(), &proto.SelectBaanRequest{UserId: t.UserMock.ID.String(), Baans: baanIds})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestUpdateBaanSelectionInvalidNumberOfBaan() {
	testUpdateBaanSelectionInvalidNumberOfBaan(t.T(), t.UserMock.ID.String(), []string{}, t.conf)
	testUpdateBaanSelectionInvalidNumberOfBaan(t.T(), t.UserMock.ID.String(), []string{"1", "2"}, t.conf)
	testUpdateBaanSelectionInvalidNumberOfBaan(t.T(), t.UserMock.ID.String(), []string{"1", "2", "3", "4"}, t.conf)
}

func testUpdateBaanSelectionInvalidNumberOfBaan(t *testing.T, userId string, baanIds []string, conf config.App) {
	repo := &mock.RepositoryMock{}

	userRepo := &mockUser.RepositoryMock{}

	fileSrv := &mockFile.ServiceMock{}

	bgsRepo := &mockBgs.RepositoryMock{}

	cacheRepo := &mockCache.RepositoryMock{}

	baanRepo := &mockBaan.RepositoryMock{}

	srv := NewService(repo, userRepo, bgsRepo, fileSrv, cacheRepo, baanRepo, conf)

	actual, err := srv.SelectBaan(context.Background(), &proto.SelectBaanRequest{UserId: userId, Baans: baanIds})

	st, ok := status.FromError(err)

	assert.True(t, ok)
	assert.Nil(t, actual)
	assert.Equal(t, codes.InvalidArgument, st.Code())
}
