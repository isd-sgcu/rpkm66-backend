package file

import (
	file_svc "github.com/isd-sgcu/rpkm66-backend/internal/service/file"
	proto "github.com/isd-sgcu/rpkm66-go-proto/rpkm66/file/file/v1"
)

type Service interface {
	GetSignedUrl(uid string) (string, error)
}

func NewService(client proto.FileServiceClient) Service {
	return file_svc.NewService(client)
}
