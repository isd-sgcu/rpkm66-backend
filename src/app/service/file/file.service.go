package file

import (
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	client proto.FileServiceClient
}

func NewService(client proto.FileServiceClient) *Service {
	return &Service{client: client}
}

func (s *Service) GetSignedUrl(uid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetSignedUrl(ctx, &proto.GetSignedUrlRequest{UserId: uid})
	if err != nil {
		return "", status.Error(codes.Unavailable, "Error while connecting to the file service")
	}

	return res.Url, nil
}
