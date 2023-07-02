package file

import (
	"time"

	proto "github.com/isd-sgcu/rpkm66-go-proto/rpkm66/file/file/v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type serviceImpl struct {
	client proto.FileServiceClient
}

func NewService(client proto.FileServiceClient) *serviceImpl {
	return &serviceImpl{client: client}
}

func (s *serviceImpl) GetSignedUrl(uid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetSignedUrl(ctx, &proto.GetSignedUrlRequest{UserId: uid})
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "file").
			Str("module", "get signed url").
			Msg("Error while connecting to service")
		return "", err
	}

	return res.Url, nil
}
