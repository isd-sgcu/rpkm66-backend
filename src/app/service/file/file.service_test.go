package file

import (
	"github.com/bxcodec/faker/v3"
	mock "github.com/isd-sgcu/rnkm65-backend/src/mocks/file"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type FileServiceTest struct {
	suite.Suite
	userId string
	url    string
}

func TestFileService(t *testing.T) {
	suite.Run(t, new(FileServiceTest))
}

func (t *FileServiceTest) SetupTest() {
	t.userId = faker.UUIDDigit()
	t.url = faker.URL()
}

func (t *FileServiceTest) TestGetSignUrlSuccess() {
	want := t.url

	c := mock.ClientMock{}
	c.On("GetSignedUrl", &proto.GetSignedUrlRequest{UserId: t.userId}).Return(&proto.GetSignedUrlResponse{Url: t.url}, nil)

	srv := NewService(&c)
	actual, err := srv.GetSignedUrl(t.userId)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}
func (t *FileServiceTest) TestGetSignUrlFailed() {
	c := mock.ClientMock{}
	c.On("GetSignedUrl", &proto.GetSignedUrlRequest{UserId: t.userId}).Return(nil, status.Error(codes.Unavailable, "Connection timeout"))

	srv := NewService(&c)
	actual, err := srv.GetSignedUrl(t.userId)

	st, ok := status.FromError(err)
	assert.True(t.T(), ok)
	assert.Equal(t.T(), "", actual)
	assert.Equal(t.T(), codes.Unavailable, st.Code())
}
