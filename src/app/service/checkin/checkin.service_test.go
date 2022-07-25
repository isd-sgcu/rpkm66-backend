package checkin

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	cst "github.com/isd-sgcu/rnkm65-backend/src/constant/checkin"
	cmock "github.com/isd-sgcu/rnkm65-backend/src/mocks/cache"
	rmock "github.com/isd-sgcu/rnkm65-backend/src/mocks/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CheckinServiceTest struct {
	suite.Suite
	Ci          *checkin.Checkin
	Token       string
	CiToken     *checkin.CheckinToken
	CiTokenInfo *checkin.TokenInfo
	Conf        config.App
}

func TestCheckinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	t.Ci = &checkin.Checkin{
		UserId:     uuid.New().String(),
		EventType:  1,
		CheckinAt:  utils.GetCurrentTimePtr(),
		CheckoutAt: utils.GetCurrentTimePtr(),
	}

	t.Token = uuid.New().String()

	t.CiToken = &checkin.CheckinToken{
		Token:       t.Token,
		UserId:      t.Ci.UserId,
		CheckinType: cst.CHECKIN,
	}

	t.CiTokenInfo = &checkin.TokenInfo{
		Id:          t.Ci.UserId,
		CheckinType: cst.CHECKIN,
		EventType:   1,
	}
}

func (t *CheckinServiceTest) TestCheckinVerifyCached() {
	want := &proto.CheckinVerifyResponse{
		CheckinToken: t.Token,
		CheckinType:  cst.CHECKIN,
	}

	req := &proto.CheckinVerifyRequest{
		Id:        t.Ci.UserId,
		EventType: t.Ci.EventType,
	}

	ciToken := &checkin.CheckinToken{}

	repo := &rmock.RepositoryMock{}

	cr := newCacheMock()
	cr.On("GetCache", utils.GetCacheKey(req.Id, req.EventType), ciToken).Return(t.CiToken, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinVerify(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), actual, want)
}

func (t *CheckinServiceTest) TestCheckinVerifyNewToken() {
	req := &proto.CheckinVerifyRequest{
		Id:        t.Ci.UserId,
		EventType: t.Ci.EventType,
	}

	repo := &rmock.RepositoryMock{}
	ciToken := &checkin.CheckinToken{}
	cr := newCacheMock()

	repo.On("FindLastCheckin", t.Ci.UserId, t.Ci.EventType, &checkin.Checkin{}).Return(nil, gorm.ErrRecordNotFound)

	cr.On("GetCache", utils.GetCacheKey(req.Id, req.EventType), ciToken).Return(nil, redis.Nil)
	cr.On("SaveCache", utils.GetCacheKey(t.Ci.UserId, req.EventType), mock.Anything, mock.Anything).Return(nil)
	cr.On("SaveCache", mock.Anything, t.CiTokenInfo, mock.Anything).Return(nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinVerify(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), len(actual.CheckinToken), 36)
}

func (t *CheckinServiceTest) TestCheckinVerifyUnknown() {
	req := &proto.CheckinVerifyRequest{
		Id:        t.Ci.UserId,
		EventType: t.Ci.EventType,
	}

	repo := &rmock.RepositoryMock{}

	ci := &checkin.Checkin{}
	repo.On("FindLastCheckin", t.Ci.UserId, t.Ci.EventType, ci).Return(nil, gorm.ErrInvalidData)

	cr := newCacheMock()
	cr.On("GetCache", utils.GetCacheKey(req.Id, req.EventType), &checkin.CheckinToken{}).Return(nil, redis.Nil)

	service := NewService(repo, cr, t.Conf)

	_, err := service.CheckinVerify(context.Background(), req)

	assert.NotNil(t.T(), err)
}

func (t *CheckinServiceTest) TestCheckinVerifyInvalidToken() {
	repo := &rmock.RepositoryMock{}

	randomUUID, _ := uuid.NewUUID()

	req := &proto.CheckinConfirmRequest{
		Token: randomUUID.String(),
	}

	cr := newCacheMock()
	cr.On("GetCache", mock.AnythingOfType("string"), &checkin.TokenInfo{}).Return(nil, redis.Nil)

	service := NewService(repo, cr, t.Conf)

	_, err := service.CheckinConfirm(context.Background(), req)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Equal(t.T(), codes.PermissionDenied, st.Code())
}

func (t *CheckinServiceTest) TestCheckinConfirmUnknown() {
	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	repo := &rmock.RepositoryMock{}
	repo.On("Checkin", newCheckin(t.Ci.UserId, t.Ci.EventType)).Return(status.Error(codes.Internal, "Internal Error"))

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(t.CiTokenInfo, nil)

	service := NewService(repo, cr, t.Conf)

	_, err := service.CheckinConfirm(context.Background(), req)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func (t *CheckinServiceTest) TestCheckinConfirmSuccess() {
	want := &proto.CheckinConfirmResponse{
		Success: true,
	}

	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	repo := &rmock.RepositoryMock{}
	repo.On("Checkin", newCheckin(t.Ci.UserId, t.Ci.EventType)).Return(nil)

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(t.CiTokenInfo, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinConfirm(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CheckinServiceTest) TestCheckoutConfirmSuccess() {
	want := &proto.CheckinConfirmResponse{
		Success: true,
	}

	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	repo := &rmock.RepositoryMock{}

	repo.On("Checkout", mock.Anything).Return(nil)

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(&checkin.TokenInfo{
		Id:          t.Ci.UserId,
		CheckinType: cst.CHECKOUT,
		EventType:   t.Ci.EventType,
	}, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinConfirm(context.Background(), req)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CheckinServiceTest) TestCheckinConfirmInvalidUserId() {
	req := &proto.CheckinConfirmRequest{
		Token: t.Token,
	}

	randomUUID := uuid.New().String()

	repo := &rmock.RepositoryMock{}

	ci := newCheckin(randomUUID, t.Ci.EventType)

	repo.On("Checkin", ci).Return(errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails"))

	cr := newCacheMock()
	cr.On("GetCache", t.Token, &checkin.TokenInfo{}).Return(&checkin.TokenInfo{
		Id:          randomUUID,
		CheckinType: cst.CHECKIN,
		EventType:   t.Ci.EventType,
	}, nil)

	service := NewService(repo, cr, t.Conf)

	actual, err := service.CheckinConfirm(context.Background(), req)

	st, ok := status.FromError(err)

	assert.True(t.T(), ok)
	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), codes.Internal, st.Code())
}

func newCacheMock() *cmock.RepositoryMock {
	return &cmock.RepositoryMock{
		V: make(map[string]interface{}),
	}
}
