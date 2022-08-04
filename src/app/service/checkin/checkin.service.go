package checkin

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/checkin"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	cst "github.com/isd-sgcu/rnkm65-backend/src/constant/checkin"
)

type IRepository interface {
	Checkin(*checkin.Checkin) error
	Checkout(*checkin.Checkin) error
	FindLastCheckin(string, int32, *checkin.Checkin) error
}

type ICacheRepository interface {
	SaveCache(string, interface{}, int) error
	GetCache(string, interface{}) error
	RemoveCache(key string) error
}

type Service struct {
	repo  IRepository
	cache ICacheRepository
	conf  config.App
}

func NewService(repo IRepository, cache ICacheRepository, conf config.App) *Service {
	return &Service{repo: repo, cache: cache, conf: conf}
}

func (s *Service) CheckinVerify(_ context.Context, req *proto.CheckinVerifyRequest) (*proto.CheckinVerifyResponse, error) {
	ciToken := &checkin.CheckinToken{}
	err := s.cache.GetCache(utils.GetCacheKey(req.Id, req.EventType), ciToken)

	if err != redis.Nil {
		return &proto.CheckinVerifyResponse{
			CheckinToken: ciToken.Token,
			CheckinType:  ciToken.CheckinType,
		}, nil
	}

	ci := &checkin.Checkin{}
	err = s.repo.FindLastCheckin(req.Id, req.EventType, ci)

	var checkinType int32

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().
			Err(err).
			Str("service", "checkin").
			Str("module", "checkinverify").
			Str("user_id", req.Id).
			Msg("Unknown db error")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	if ci.CheckinAt != nil && ci.CheckoutAt == nil {
		checkinType = cst.CHECKOUT
	} else {
		checkinType = cst.CHECKIN
	}

	_token, err := uuid.NewUUID()

	if err != nil {
		log.Error().
			Err(err).
			Str("service", "checkin").
			Str("module", "checkinverify").
			Str("user_id", req.Id).
			Msg("UUID broken")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	token := _token.String()

	err = s.cache.SaveCache(token, &checkin.TokenInfo{
		Id:          req.Id,
		CheckinType: checkinType,
		EventType:   req.EventType,
	}, s.conf.CICacheTTL)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "checkin").
			Str("module", "checkinverify").
			Str("user_id", req.Id).
			Msg("Error while connecting to redis server")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	err = s.cache.SaveCache(utils.GetCacheKey(req.Id, req.EventType), &checkin.CheckinToken{
		Token:       token,
		UserId:      req.Id,
		CheckinType: checkinType,
	}, s.conf.CICacheTTL)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "checkin").
			Str("module", "checkinverify").
			Str("user_id", req.Id).
			Msg("Error while connecting to redis server")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	res := &proto.CheckinVerifyResponse{
		CheckinToken: token,
		CheckinType:  checkinType,
	}

	log.Info().
		Str("service", "checkin").
		Str("module", "checkinverify").
		Str("user_id", req.Id).
		Msg("Successfully send verification checkin to user")

	return res, nil
}

func (s *Service) CheckinConfirm(_ context.Context, req *proto.CheckinConfirmRequest) (*proto.CheckinConfirmResponse, error) {
	cached := &checkin.TokenInfo{}

	err := s.cache.GetCache(req.Token, cached)

	if err == redis.Nil {
		return nil, status.Error(codes.PermissionDenied, "Invalid token")
	}

	defer func() {
		if err := s.cache.RemoveCache(utils.GetCacheKey(cached.Id, cached.EventType)); err != nil {
			log.Error().Err(err).
				Str("service", "Checkin").
				Str("module", "checkin confirm").
				Str("user_id", cached.Id).
				Msg("Error while removing user cache")
		}
	}()

	defer func() {
		if err := s.cache.RemoveCache(req.Token); err != nil {
			log.Error().Err(err).
				Str("service", "Checkin").
				Str("module", "checkin confirm").
				Str("user_id", cached.Id).
				Msg("Error while removing token cache")
		}
	}()

	switch cached.CheckinType {
	case cst.CHECKIN:
		ci := newCheckin(cached.Id, cached.EventType)
		err = s.repo.Checkin(ci)
	case cst.CHECKOUT:
		ci := newCheckout(cached.Id, cached.EventType)
		err = s.repo.Checkout(ci)
	default:
		return nil, status.Error(codes.InvalidArgument, "Invalid checkin type")
	}

	if err != nil {
		log.Error().Err(err).
			Str("service", "Checkin").
			Str("module", "checkin confirm").
			Str("user_id", cached.Id).
			Msg("Error while Checkin, Possibly due to invalid user-uuid")

		return nil, status.Error(codes.Internal, "Internal Error")
	}

	log.Info().
		Str("service", "checkin").
		Str("module", "checkin confirm").
		Str("user_id", cached.Id).
		Msg("Successfully checkin user")

	return &proto.CheckinConfirmResponse{
		Success: true,
	}, nil
}

func newCheckin(userid string, eventType int32) *checkin.Checkin {
	return &checkin.Checkin{
		UserId:    userid,
		EventType: eventType,
	}
}

func newCheckout(userid string, eventType int32) *checkin.Checkin {
	return &checkin.Checkin{
		UserId:     userid,
		CheckoutAt: utils.GetCurrentTimePtr(),
		EventType:  eventType,
	}
}
