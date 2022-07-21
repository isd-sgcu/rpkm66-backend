package baan

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	"github.com/isd-sgcu/rnkm65-backend/src/config"
	constant "github.com/isd-sgcu/rnkm65-backend/src/constant/baan"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repository IRepository
	cache      ICacheRepository
	conf       config.App
}

type IRepository interface {
	FindAll(*[]*baan.Baan) error
	FindOne(string, *baan.Baan) error
}

type ICacheRepository interface {
	SaveCache(string, interface{}, int) error
	GetCache(string, interface{}) error
}

func NewService(repository IRepository, cache ICacheRepository, conf config.App) *Service {
	return &Service{
		repository: repository,
		cache:      cache,
		conf:       conf,
	}
}

func (s *Service) FindAllBaan(_ context.Context, _ *proto.FindAllBaanRequest) (*proto.FindAllBaanResponse, error) {
	var baans []*baan.Baan
	err := s.cache.GetCache(constant.BaanKey, &baans)
	if err != redis.Nil {

		if err != nil {
			log.Error().
				Err(err).
				Str("service", "baan").
				Str("module", "find all").
				Msg("Error while get cache")

			return nil, status.Error(codes.Unavailable, "Service is down")
		}

		return &proto.FindAllBaanResponse{Baans: RawToDtoList(&baans)}, nil
	}

	err = s.repository.FindAll(&baans)
	if err != nil {

		log.Error().Err(err).
			Str("service", "baan").
			Str("module", "find all").
			Msg("Error while querying all baans")

		return nil, status.Error(codes.Unavailable, "Internal error")
	}

	err = s.cache.SaveCache(constant.BaanKey, &baans, s.conf.BaanCacheTTL)
	if err != nil {

		log.Error().
			Err(err).
			Str("service", "baan").
			Str("module", "find all").
			Msg("Error while saving the cache")

		return nil, status.Error(codes.Unavailable, "Service is down")
	}

	return &proto.FindAllBaanResponse{Baans: RawToDtoList(&baans)}, nil
}

func (s *Service) FindOneBaan(_ context.Context, req *proto.FindOneBaanRequest) (*proto.FindOneBaanResponse, error) {
	result := baan.Baan{}
	err := s.cache.GetCache(req.Id, &result)
	if err != redis.Nil {

		if err != nil {
			log.Error().
				Err(err).
				Str("service", "baan").
				Str("module", "find one").
				Msg("Error while get cache")

			return nil, status.Error(codes.Unavailable, "Error while connecting to cache server")
		}

		return &proto.FindOneBaanResponse{Baan: RawToDto(&result)}, nil
	}

	err = s.repository.FindOne(req.Id, &result)
	if err != nil {

		log.Error().
			Err(err).
			Str("service", "baan").
			Str("module", "find one").
			Msg("Not found baan")

		return nil, status.Error(codes.NotFound, "Not found baan")
	}

	err = s.cache.SaveCache(result.ID.String(), &result, s.conf.BaanCacheTTL)
	if err != nil {

		log.Error().
			Err(err).
			Str("service", "baan").
			Str("module", "find all").
			Msg("Error while saving the cache")

		return nil, status.Error(codes.Unavailable, "Error while connecting to cache server")
	}

	return &proto.FindOneBaanResponse{Baan: RawToDto(&result)}, nil
}

func RawToDtoInfoList(in *[]*baan.Baan) []*proto.BaanInfo {
	var result []*proto.BaanInfo
	for _, b := range *in {
		bi := &proto.BaanInfo{
			Id:       b.ID.String(),
			NameTH:   b.NameTH,
			NameEN:   b.NameEN,
			ImageUrl: b.ImageUrl,
		}

		result = append(result, bi)
	}

	return result
}

func RawToDtoList(in *[]*baan.Baan) []*proto.Baan {
	var result []*proto.Baan
	for _, b := range *in {
		result = append(result, RawToDto(b))
	}

	return result
}

func RawToDto(in *baan.Baan) *proto.Baan {
	return &proto.Baan{
		Id:            in.ID.String(),
		NameTH:        in.NameTH,
		DescriptionTH: in.DescriptionTH,
		NameEN:        in.NameEN,
		DescriptionEN: in.DescriptionEN,
		Size:          proto.BaanSize(in.Size),
		Facebook:      in.Facebook,
		FacebookUrl:   in.FacebookUrl,
		Instagram:     in.Instagram,
		InstagramUrl:  in.InstagramUrl,
		Line:          in.Line,
		LineUrl:       in.LineUrl,
		ImageUrl:      in.ImageUrl,
	}
}
