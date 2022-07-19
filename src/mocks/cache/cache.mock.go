package cache

import (
	dto "github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
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
		}
	}

	return args.Error(1)
}
