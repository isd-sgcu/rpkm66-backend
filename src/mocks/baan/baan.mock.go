package baan

import (
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) FindAll(in *[]*baan.Baan) error {
	args := r.Called(*in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*[]*baan.Baan)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindOne(id string, in *baan.Baan) error {
	args := r.Called(id, in)

	if args.Get(0) != nil {
		*in = *args.Get(0).(*baan.Baan)
	}

	return args.Error(1)
}

func (r *RepositoryMock) FindMulti(ids []string, in *[]*baan.Baan) error {
	args := r.Called(ids, in)

	if args.Get(0) != nil {
		*in = args.Get(0).([]*baan.Baan)
	}

	return args.Error(1)
}
