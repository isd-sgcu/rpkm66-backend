package file

import (
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
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
		log.Error().
			Err(err).
			Str("service", "file").
			Str("module", "get signed url").
			Msg("Error while connecting to service")
		return "", err
	}

	return res.Url, nil
}
