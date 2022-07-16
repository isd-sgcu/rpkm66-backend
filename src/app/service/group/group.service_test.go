package group

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/group"
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

type GroupServiceTest struct {
	suite.Suite
	Group              *group.Group
	GroupDto           *proto.Group
	CreateGroupReqMock *proto.CreateGroupRequest
	UpdateGroupReqMock *proto.UpdateGroupRequest
	UserMock           *user.User
	UserDtoMock        *proto.User
}

func TestUserService(t *testing.T) {
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
		GroupID:         utils.UUIDAdr(uuid.New()),
	}
	t.UserDtoMock = &proto.User{
		Id:              t.UserMock.ID.String(),
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
		CanSelectBaan:   *t.UserMock.CanSelectBaan,
		GroupId:         t.UserMock.GroupID.String(),
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

	t.GroupDto = &proto.Group{
		Id:       t.Group.ID.String(),
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*proto.User{t.UserDtoMock},
	}

	t.CreateGroupReqMock = &proto.CreateGroupRequest{
		UserId: t.UserMock.ID.String(),
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
}

func (t *GroupServiceTest) TestFindByTokenSuccess() {
	want := &proto.FindByTokenGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)

	srv := NewService(repo)
	actual, err := srv.FindByToken(context.Background(), &proto.FindByTokenGroupRequest{Token: t.GroupDto.Token})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestFindByTokenNotFound() {
	repo := &mock.RepositoryMock{}

	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.FindByToken(context.Background(), &proto.FindByTokenGroupRequest{Token: t.GroupDto.Token})

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestCreateSuccess() {
	want := &proto.CreateGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	in := &group.Group{
		LeaderID: t.UserMock.ID.String(),
	}

	usr := &user.User{}

	repo.On("FindUserById", t.UserMock.ID.String(), usr).Return(t.UserMock, nil)
	repo.On("Create", in).Return(t.Group, nil)
	repo.On("UpdateUser", t.UserMock).Return(t.UserMock, nil)

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateGroupReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestCreateNotFound() {
	usr := &user.User{}

	repo := &mock.RepositoryMock{}
	repo.On("FindUserById", t.Group.LeaderID, usr).Return(nil, errors.New("not found user"))

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), &proto.CreateGroupRequest{UserId: t.Group.LeaderID})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestCreateMalformed() {
	repo := &mock.RepositoryMock{}

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), &proto.CreateGroupRequest{UserId: "abc"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestCreateInternalErr() {
	repo := &mock.RepositoryMock{}

	in := &group.Group{
		LeaderID: t.Group.LeaderID,
	}

	usr := &user.User{}
	repo.On("FindUserById", t.Group.LeaderID, usr).Return(t.UserMock, nil)
	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	srv := NewService(repo)

	actual, err := srv.Create(context.Background(), t.CreateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *GroupServiceTest) TestUpdateSuccess() {
	want := &proto.UpdateGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}

	repo.On("UpdateWithLeader", t.Group.LeaderID, t.Group).Return(t.Group, nil)

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestUpdateNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("UpdateWithLeader", t.Group.LeaderID, t.Group).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestUpdateMalformed() {
	repo := &mock.RepositoryMock{}
	repo.On("UpdateWithLeader", t.Group.LeaderID, t.Group).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)

	t.UpdateGroupReqMock.Group.Id = "abc"

	actual, err := srv.Update(context.Background(), t.UpdateGroupReqMock)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

//Case1 : a user is a not a king in the group --> expected result : the user is able to join other group, and remove the user from the previous group
func (t *GroupServiceTest) TestJoinSuccess1() {
	beforeJoinedUser := &user.User{
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
		GroupID:         utils.UUIDAdr(uuid.New()),
	}
	afterJoinedUser := &user.User{
		Base: model.Base{
			ID:        beforeJoinedUser.ID,
			CreatedAt: beforeJoinedUser.CreatedAt,
			UpdatedAt: beforeJoinedUser.UpdatedAt,
			DeletedAt: beforeJoinedUser.DeletedAt,
		},
		Title:           beforeJoinedUser.Title,
		Firstname:       beforeJoinedUser.Firstname,
		Lastname:        beforeJoinedUser.Lastname,
		Nickname:        beforeJoinedUser.Nickname,
		StudentID:       beforeJoinedUser.StudentID,
		Faculty:         beforeJoinedUser.Faculty,
		Year:            beforeJoinedUser.Year,
		Phone:           beforeJoinedUser.Phone,
		LineID:          beforeJoinedUser.LineID,
		Email:           beforeJoinedUser.Email,
		AllergyFood:     beforeJoinedUser.AllergyFood,
		FoodRestriction: beforeJoinedUser.FoodRestriction,
		AllergyMedicine: beforeJoinedUser.AllergyMedicine,
		Disease:         beforeJoinedUser.Disease,
		CanSelectBaan:   beforeJoinedUser.CanSelectBaan,
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}
	afterJoinedUserDto := &proto.User{
		Id:              afterJoinedUser.ID.String(),
		Title:           afterJoinedUser.Title,
		Firstname:       afterJoinedUser.Firstname,
		Lastname:        afterJoinedUser.Lastname,
		Nickname:        afterJoinedUser.Nickname,
		StudentID:       afterJoinedUser.StudentID,
		Faculty:         afterJoinedUser.Faculty,
		Year:            afterJoinedUser.Year,
		Phone:           afterJoinedUser.Phone,
		LineID:          afterJoinedUser.LineID,
		Email:           afterJoinedUser.Email,
		AllergyFood:     afterJoinedUser.AllergyFood,
		FoodRestriction: afterJoinedUser.FoodRestriction,
		AllergyMedicine: afterJoinedUser.AllergyMedicine,
		Disease:         afterJoinedUser.Disease,
		CanSelectBaan:   *afterJoinedUser.CanSelectBaan,
		GroupId:         afterJoinedUser.GroupID.String(),
	}

	//joinedGrp := &group.Group{
	//	Base: model.Base{
	//		ID:        t.Group.ID,
	//		CreatedAt: t.Group.CreatedAt,
	//		UpdatedAt: t.Group.UpdatedAt,
	//		DeletedAt: t.Group.DeletedAt,
	//	},
	//	LeaderID: t.Group.LeaderID,
	//	Token:    t.Group.Token,
	//	Members:  []*user.User{t.UserMock, afterJoinedUser},
	//}

	want := &proto.JoinGroupResponse{Group: &proto.Group{
		Id:       t.Group.ID.String(),
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*proto.User{t.UserDtoMock, afterJoinedUserDto},
	}}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)
	repo.On("FindUserById", beforeJoinedUser.ID.String(), &user.User{}).Return(beforeJoinedUser, nil)
	repo.On("UpdateUser", afterJoinedUser).Return(afterJoinedUser, nil)

	srv := NewService(repo)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.GroupDto.Token, UserId: beforeJoinedUser.ID.String(), IsLeader: false, Members: 2})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

//Case2 : a user is only one in the group --> expected result : the user is able to join other group, and delete the previous group
func (t *GroupServiceTest) TestJoinSuccess2() {
	headUser := &user.User{
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
		GroupID:         utils.UUIDAdr(uuid.New()),
	}
	headUserDto := &proto.User{
		Id:              headUser.ID.String(),
		Title:           headUser.Title,
		Firstname:       headUser.Firstname,
		Lastname:        headUser.Lastname,
		Nickname:        headUser.Nickname,
		StudentID:       headUser.StudentID,
		Faculty:         headUser.Faculty,
		Year:            headUser.Year,
		Phone:           headUser.Phone,
		LineID:          headUser.LineID,
		Email:           headUser.Email,
		AllergyFood:     headUser.AllergyFood,
		FoodRestriction: headUser.FoodRestriction,
		AllergyMedicine: headUser.AllergyMedicine,
		Disease:         headUser.Disease,
		CanSelectBaan:   *headUser.CanSelectBaan,
		GroupId:         headUser.GroupID.String(),
	}
	prevGrp := &group.Group{
		Base: model.Base{
			ID:        *headUser.GroupID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: headUser.ID.String(),
		Token:    faker.Word(),
		Members:  []*user.User{headUser},
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
		CanSelectBaan:   utils.BoolAdr(true),
		GroupID:         utils.UUIDAdr(prevGrp.ID),
	}
	joinUserDto := &proto.User{
		Id:              joinUser.ID.String(),
		Title:           joinUser.Title,
		Firstname:       joinUser.Firstname,
		Lastname:        joinUser.Lastname,
		Nickname:        joinUser.Nickname,
		StudentID:       joinUser.StudentID,
		Faculty:         joinUser.Faculty,
		Year:            joinUser.Year,
		Phone:           joinUser.Phone,
		LineID:          joinUser.LineID,
		Email:           joinUser.Email,
		AllergyFood:     joinUser.AllergyFood,
		FoodRestriction: joinUser.FoodRestriction,
		AllergyMedicine: joinUser.AllergyMedicine,
		Disease:         joinUser.Disease,
		CanSelectBaan:   *joinUser.CanSelectBaan,
		GroupId:         joinUser.GroupID.String(),
	}

	want := &proto.JoinGroupResponse{Group: &proto.Group{
		Id:       headUserDto.GroupId,
		LeaderID: headUserDto.Id,
		Token:    prevGrp.Token,
		Members:  []*proto.User{headUserDto, joinUserDto},
	}}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", prevGrp.Token, &group.Group{}).Return(prevGrp, nil)
	repo.On("FindUserById", t.UserDtoMock.Id, &user.User{}).Return(t.UserMock, nil)
	repo.On("UpdateUser", joinUser).Return(joinUser, nil)
	repo.On("Delete", t.UserDtoMock.GroupId).Return(nil)

	srv := NewService(repo)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: prevGrp.Token, UserId: t.UserDtoMock.Id, IsLeader: true, Members: 1})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

//Case3 : a leader of the group with members can not join any group
func (t *GroupServiceTest) TestJoinForbidden() {
	repo := &mock.RepositoryMock{}

	srv := NewService(repo)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: faker.Word(), UserId: t.UserDtoMock.Id, IsLeader: true, Members: 3})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

// Wrong Token
func (t *GroupServiceTest) TestJoinNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.GroupDto.Token, UserId: uuid.New().String(), IsLeader: false, Members: 2})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

//Wrong userId
func (t *GroupServiceTest) TestJoinMalformed() {
	repo := &mock.RepositoryMock{}

	srv := NewService(repo)

	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: t.GroupDto.Token, UserId: "abc", IsLeader: false, Members: 2})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestJoinFullGroup() {
	headUser := &user.User{
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
		GroupID:         utils.UUIDAdr(uuid.New()),
	}
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
		GroupID:         headUser.GroupID,
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
		GroupID:         headUser.GroupID,
	}

	fullGroup := &group.Group{
		Base: model.Base{
			ID:        *headUser.GroupID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: headUser.ID.String(),
		Token:    faker.Word(),
		Members:  []*user.User{headUser, member1User, member2User},
	}
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", fullGroup.Token, &group.Group{}).Return(fullGroup, nil)

	srv := NewService(repo)

	actual, err := srv.Join(context.Background(), &proto.JoinGroupRequest{Token: fullGroup.Token, UserId: t.UserDtoMock.Id, IsLeader: true, Members: 1})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *GroupServiceTest) TestDeleteMemberSuccess() {
	deletedUser := &user.User{
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
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}

	prevGroup := &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{t.UserMock, deletedUser},
	}

	in := &group.Group{
		LeaderID: deletedUser.ID.String(),
	}
	createdUser := &user.User{
		Base: model.Base{
			ID:        deletedUser.ID,
			CreatedAt: deletedUser.CreatedAt,
			UpdatedAt: deletedUser.UpdatedAt,
			DeletedAt: deletedUser.DeletedAt,
		},
		Title:           deletedUser.Title,
		Firstname:       deletedUser.Firstname,
		Lastname:        deletedUser.Lastname,
		Nickname:        deletedUser.Nickname,
		StudentID:       deletedUser.StudentID,
		Faculty:         deletedUser.Faculty,
		Year:            deletedUser.Year,
		Phone:           deletedUser.Phone,
		LineID:          deletedUser.LineID,
		Email:           deletedUser.Email,
		AllergyFood:     deletedUser.AllergyFood,
		FoodRestriction: deletedUser.FoodRestriction,
		AllergyMedicine: deletedUser.AllergyMedicine,
		Disease:         deletedUser.Disease,
		CanSelectBaan:   utils.BoolAdr(true),
		GroupID:         utils.UUIDAdr(in.ID),
	}
	want := &proto.DeleteMemberGroupResponse{Group: t.GroupDto}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.GroupDto.Token, &group.Group{}).Return(prevGroup, nil)
	repo.On("FindUserById", deletedUser.ID.String(), &user.User{}).Return(deletedUser, nil)
	repo.On("Create", in).Return(in, nil)
	repo.On("UpdateUser", createdUser).Return(createdUser, nil)

	srv := NewService(repo)
	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{Token: t.GroupDto.Token, DeletedId: deletedUser.ID.String(), UserId: t.GroupDto.LeaderID})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestDeleteMemberForbidden() {
	deletedUser := &user.User{
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
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}

	prevGroup := &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{t.UserMock, deletedUser},
	}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.GroupDto.Token, &group.Group{}).Return(prevGroup, nil)

	srv := NewService(repo)
	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{Token: t.GroupDto.Token, DeletedId: deletedUser.ID.String(), UserId: deletedUser.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *GroupServiceTest) TestDeleteMemberNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.GroupDto.Token, &group.Group{}).Return(nil, errors.New("Not found group"))

	srv := NewService(repo)
	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{Token: t.GroupDto.Token, DeletedId: uuid.New().String(), UserId: t.GroupDto.LeaderID})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestDeleteMemberMalformed() {
	repo := &mock.RepositoryMock{}

	srv := NewService(repo)

	actual, err := srv.DeleteMember(context.Background(), &proto.DeleteMemberGroupRequest{Token: t.GroupDto.Token, DeletedId: "abc", UserId: t.GroupDto.LeaderID})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupSuccess() {
	leavedUser := &user.User{
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
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}
	prevGrp := &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{t.UserMock, leavedUser},
	}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", prevGrp.Token, &group.Group{}).Return(prevGrp, nil)

	in := &group.Group{
		LeaderID: leavedUser.ID.String(),
	}
	repo.On("Create", in).Return(in, nil)

	updatedUser := &user.User{
		Base: model.Base{
			ID:        leavedUser.ID,
			CreatedAt: leavedUser.CreatedAt,
			UpdatedAt: leavedUser.UpdatedAt,
			DeletedAt: leavedUser.DeletedAt,
		},
		Title:           leavedUser.Title,
		Firstname:       leavedUser.Firstname,
		Lastname:        leavedUser.Lastname,
		Nickname:        leavedUser.Nickname,
		StudentID:       leavedUser.StudentID,
		Faculty:         leavedUser.Faculty,
		Year:            leavedUser.Year,
		Phone:           leavedUser.Phone,
		LineID:          leavedUser.LineID,
		Email:           leavedUser.Email,
		AllergyFood:     leavedUser.AllergyFood,
		FoodRestriction: leavedUser.FoodRestriction,
		AllergyMedicine: leavedUser.AllergyMedicine,
		Disease:         leavedUser.Disease,
		CanSelectBaan:   leavedUser.CanSelectBaan,
		GroupID:         utils.UUIDAdr(in.ID),
	}
	updatedUserDto := &proto.User{
		Id:              updatedUser.ID.String(),
		Title:           updatedUser.Title,
		Firstname:       updatedUser.Firstname,
		Lastname:        updatedUser.Lastname,
		Nickname:        updatedUser.Nickname,
		StudentID:       updatedUser.StudentID,
		Faculty:         updatedUser.Faculty,
		Year:            updatedUser.Year,
		Phone:           updatedUser.Phone,
		LineID:          updatedUser.LineID,
		Email:           updatedUser.Email,
		AllergyFood:     updatedUser.AllergyFood,
		FoodRestriction: updatedUser.FoodRestriction,
		AllergyMedicine: updatedUser.AllergyMedicine,
		Disease:         updatedUser.Disease,
		CanSelectBaan:   *updatedUser.CanSelectBaan,
		GroupId:         updatedUser.GroupID.String(),
	}
	repo.On("UpdateUser", updatedUser).Return(updatedUser, nil)

	want := &proto.LeaveGroupResponse{
		Group: &proto.Group{
			Id:       in.ID.String(),
			LeaderID: leavedUser.ID.String(),
			Token:    in.Token,
			Members:  []*proto.User{updatedUserDto},
		},
	}

	srv := NewService(repo)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{Token: prevGrp.Token, UserId: leavedUser.ID.String()})

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *GroupServiceTest) TestLeaveGroupNotFound() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(nil, errors.New("not found group"))

	srv := NewService(repo)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{Token: t.Group.Token, UserId: uuid.New().String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.NotFound, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupMalformed() {
	repo := &mock.RepositoryMock{}

	srv := NewService(repo)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{Token: t.Group.Token, UserId: "abc"})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.InvalidArgument, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupForbidden() {
	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", t.Group.Token, &group.Group{}).Return(t.Group, nil)

	srv := NewService(repo)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{Token: t.Group.Token, UserId: t.UserMock.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *GroupServiceTest) TestLeaveGroupInternalErr() {
	leavedUser := &user.User{
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
		GroupID:         utils.UUIDAdr(t.Group.ID),
	}
	prevGrp := &group.Group{
		Base: model.Base{
			ID:        t.Group.ID,
			CreatedAt: t.Group.CreatedAt,
			UpdatedAt: t.Group.UpdatedAt,
			DeletedAt: t.Group.DeletedAt,
		},
		LeaderID: t.Group.LeaderID,
		Token:    t.Group.Token,
		Members:  []*user.User{t.UserMock, leavedUser},
	}

	repo := &mock.RepositoryMock{}
	repo.On("FindGroupByToken", prevGrp.Token, &group.Group{}).Return(prevGrp, nil)

	in := &group.Group{
		LeaderID: leavedUser.ID.String(),
	}
	repo.On("Create", in).Return(nil, errors.New("something wrong"))

	srv := NewService(repo)

	actual, err := srv.Leave(context.Background(), &proto.LeaveGroupRequest{Token: prevGrp.Token, UserId: leavedUser.ID.String()})

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}
