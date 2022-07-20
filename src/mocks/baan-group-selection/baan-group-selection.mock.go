package baan_group_selection

import (
	baan_group_selection "github.com/isd-sgcu/rnkm65-backend/src/app/model/baan-group-selection"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) SaveBaansSelection(in *[]*baan_group_selection.BaanGroupSelection) error {
	args := r.Called(in)

	if args.Get(0) != nil {
		*in = args.Get(0).([]*baan_group_selection.BaanGroupSelection)
	}

	return args.Error(1)
}
