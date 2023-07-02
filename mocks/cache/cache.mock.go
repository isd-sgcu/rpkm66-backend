package cache

import (
	dto "github.com/isd-sgcu/rpkm66-backend/internal/entity/baan"
	"github.com/isd-sgcu/rpkm66-backend/internal/entity/checkin"
	proto "github.com/isd-sgcu/rpkm66-backend/internal/proto/rpkm66/backend/baan/v1"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
	V map[string]interface{}
}

func (t *RepositoryMock) SaveCache(key string, v interface{}, ttl int) error {
	args := t.Called(key, v, ttl)

	t.V[key] = v

	return args.Error(0)
}

func (t *RepositoryMock) GetCache(key string, v interface{}) error {
	args := t.Called(key, v)

	if args.Get(0) != nil {
		switch args.Get(0).(type) {
		case *[]*dto.Baan:
			*v.(*[]*dto.Baan) = *args.Get(0).(*[]*dto.Baan)
		case *dto.Baan:
			*v.(*dto.Baan) = *args.Get(0).(*dto.Baan)
		case *[]*proto.BaanInfo:
			*v.(*[]*proto.BaanInfo) = *args.Get(0).(*[]*proto.BaanInfo)
		case *checkin.CheckinToken:
			*v.(*checkin.CheckinToken) = *args.Get(0).(*checkin.CheckinToken)
		case *checkin.TokenInfo:
			*v.(*checkin.TokenInfo) = *args.Get(0).(*checkin.TokenInfo)
		}
	}

	return args.Error(1)
}

func (t *RepositoryMock) RemoveCache(key string) (err error) {
	delete(t.V, key)
	return err
}
