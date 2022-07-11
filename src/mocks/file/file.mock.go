package file

import (
	"context"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) GetSignedUrl(_ context.Context, in *proto.GetSignedUrlRequest, _ ...grpc.CallOption) (res *proto.GetSignedUrlResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.GetSignedUrlResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) GetSignedUrl(uid string) (res string, err error) {
	args := s.Called(uid)

	if args.Get(0) != nil {
		res = args.String(0)
	}

	return res, args.Error(1)
}
