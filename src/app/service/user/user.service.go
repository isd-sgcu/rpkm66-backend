package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/event"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/user"
	eventSrv "github.com/isd-sgcu/rnkm65-backend/src/app/service/event"
	"github.com/isd-sgcu/rnkm65-backend/src/app/utils"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	repo      IRepository
	eventRepo IEventRepository
	fileSrv   IFileService
}

type IRepository interface {
	FindOne(string, *user.User) error
	FindByStudentID(string, *user.User) error
	Create(*user.User) error
	Update(string, *user.User) error
	Verify(string, string) error
	Delete(string) error
	CreateOrUpdate(*user.User) error
	ConfirmEstamp(string, *user.User, *event.Event) error
	GetUserEstamp(string, *user.User, *[]*event.Event) error
}

type IEventRepository interface {
	FindEventByID(id string, result *event.Event) error
	FindAllEvent(result *[]*event.Event) error
}

type IFileService interface {
	GetSignedUrl(string) (string, error)
}

func NewService(repo IRepository, fileSrv IFileService, eventRepo IEventRepository) *Service {
	return &Service{repo: repo, fileSrv: fileSrv, eventRepo: eventRepo}
}

func (s *Service) FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.FindOneUserResponse, err error) {
	raw := user.User{}

	err = s.repo.FindOne(req.Id, &raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	url, err := s.fileSrv.GetSignedUrl(req.Id)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return &proto.FindOneUserResponse{User: RawToDto(&raw, "")}, nil

			case codes.Unavailable:
				log.Error().
					Err(err).
					Str("service", "user").
					Str("module", "find one").
					Str("id", req.Id).
					Msg("Something wrong")
				return nil, err

			default:
				log.Error().
					Err(err).
					Str("service", "user").
					Str("module", "find one").
					Str("id", req.Id).
					Msg("Error while connecting to service")
				return nil, err
			}
		}

		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "find one").
			Str("id", req.Id).
			Msg("Error while connecting to service")

		return nil, err
	}

	return &proto.FindOneUserResponse{User: RawToDto(&raw, url)}, nil
}

func (s *Service) FindByStudentID(_ context.Context, req *proto.FindByStudentIDUserRequest) (res *proto.FindByStudentIDUserResponse, err error) {
	result := user.User{}

	err = s.repo.FindByStudentID(req.StudentId, &result)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "find one").
			Str("student_id", req.StudentId).
			Msg("Not found user image")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res = &proto.FindByStudentIDUserResponse{User: RawToDto(&result, "")}

	return res, err
}

func (s *Service) Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.CreateUserResponse, err error) {
	raw, _ := DtoToRaw(req.User)

	err = s.repo.Create(raw)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &proto.CreateUserResponse{User: RawToDto(raw, "")}, nil
}

func (s *Service) CreateOrUpdate(_ context.Context, req *proto.CreateOrUpdateUserRequest) (result *proto.CreateOrUpdateUserResponse, err error) {
	raw, err := DtoToRaw(req.User)

	if err != nil {
		log.Error().Err(err).
			Str("module", "create or update").
			Msg("Error while mapping dto to raw")
		return nil, status.Error(codes.InvalidArgument, "invalid user id")
	}

	err = s.repo.CreateOrUpdate(raw)
	if err != nil {
		log.Error().Err(err).
			Str("module", "create or update").
			Str("student_id", raw.StudentID).
			Msg("Error while create or update the user")
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	log.Info().
		Str("service", "user").
		Str("module", "create or update").
		Str("student_id", raw.StudentID).
		Msg("Successfully create or update the user")

	return &proto.CreateOrUpdateUserResponse{User: RawToDto(raw, "")}, nil
}

func (s *Service) Verify(_ context.Context, req *proto.VerifyUserRequest) (res *proto.VerifyUserResponse, err error) {
	verify := GetColumnName(req.VerifyType)

	if verify == "" {
		log.Error().Err(err).
			Str("service", "user").
			Str("module", "verify").
			Str("type", req.VerifyType).
			Str("student_id", req.StudentId).
			Msgf("invalid type")
		return nil, status.Error(codes.InvalidArgument, "invalid type")
	}

	err = s.repo.Verify(req.StudentId, verify)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			log.Error().Err(err).
				Str("service", "user").
				Str("module", "verify").
				Str("type", req.VerifyType).
				Str("student_id", req.StudentId).
				Msgf("Cannot verify (not found)")
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Error().Err(err).
			Str("service", "user").
			Str("module", "verify").
			Str("type", req.VerifyType).
			Str("student_id", req.StudentId).
			Msgf("Cannot verify")
		return nil, status.Error(codes.Internal, "something wrong")
	}

	log.Info().
		Str("service", "user").
		Str("module", "create or update").
		Str("type", req.VerifyType).
		Str("student_id", req.StudentId).
		Msg("Successfully verify user")

	return &proto.VerifyUserResponse{Success: true}, nil
}

func (s *Service) Update(_ context.Context, req *proto.UpdateUserRequest) (res *proto.UpdateUserResponse, err error) {
	raw := &user.User{
		Title:           req.Title,
		Firstname:       req.Firstname,
		Lastname:        req.Lastname,
		Nickname:        req.Nickname,
		Phone:           req.Phone,
		LineID:          req.LineID,
		Email:           req.Email,
		AllergyFood:     req.AllergyFood,
		FoodRestriction: req.FoodRestriction,
		AllergyMedicine: req.AllergyMedicine,
		Disease:         req.Disease,
	}

	err = s.repo.Update(req.Id, raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &proto.UpdateUserResponse{User: RawToDto(raw, "")}, nil
}

func (s *Service) Delete(_ context.Context, req *proto.DeleteUserRequest) (res *proto.DeleteUserResponse, err error) {
	err = s.repo.Delete(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "something wrong")
	}

	return &proto.DeleteUserResponse{Success: true}, nil
}

func (s *Service) ConfirmEstamp(_ context.Context, req *proto.ConfirmEstampRequest) (res *proto.ConfirmEstampResponse, err error) {
	var event event.Event

	uid, err := uuid.Parse(req.UId)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "confirm estamp").
			Str("user_id", req.UId).
			Msg("Invalid id")

		return nil, status.Error(codes.InvalidArgument, "Invalid id")
	}

	user := user.User{
		Base: model.Base{
			ID: uid,
		},
	}

	err = s.eventRepo.FindEventByID(req.EId, &event)
	if err != nil {

		if errors.Is(gorm.ErrRecordNotFound, err) {
			log.Error().
				Err(err).
				Str("service", "user").
				Str("module", "confirm estamp").
				Str("user_id", req.UId).
				Msg("Not found event")

			return nil, status.Error(codes.NotFound, "Not found event")
		}

		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "confirm estamp").
			Str("user_id", req.UId).
			Msg("Error while connecting to database")

		return nil, status.Error(codes.Internal, "something wrong")
	}

	err = s.repo.ConfirmEstamp(req.UId, &user, &event)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			log.Error().
				Err(err).
				Str("service", "user").
				Str("module", "confirm estamp").
				Str("user_id", req.UId).
				Msg("Not found eStamp")

			return nil, status.Error(codes.NotFound, "Not found estamp")
		}

		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "confirm estamp").
			Str("user_id", req.UId).
			Msg("Error while connecting to database")

		return nil, status.Error(codes.Internal, "something wrong")
	}

	log.Info().
		Str("service", "user").
		Str("module", "confirm estamp").
		Str("user_id", req.UId).
		Msg("Successfully verify user")

	return &proto.ConfirmEstampResponse{}, nil
}

func (s *Service) GetUserEstamp(_ context.Context, req *proto.GetUserEstampRequest) (res *proto.GetUserEstampResponse, err error) {
	uid, err := uuid.Parse(req.UId)
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "confirm estamp").
			Str("user_id", req.UId).
			Msg("Invalid id")

		return nil, status.Error(codes.InvalidArgument, "Invalid id")
	}

	user := user.User{
		Base: model.Base{
			ID: uid,
		},
	}

	var events []*event.Event

	err = s.repo.GetUserEstamp(req.UId, &user, &events)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			log.Error().
				Err(err).
				Str("service", "user").
				Str("module", "get user estamp").
				Str("user_id", req.UId).
				Msg("Not found user")

			return nil, status.Error(codes.NotFound, "Not found user")
		}

		log.Error().
			Err(err).
			Str("service", "user").
			Str("module", "get user estamp").
			Str("user_id", req.UId).
			Msg("Error while connecting to database")

		return nil, status.Error(codes.Internal, "something wrong")
	}

	log.Info().
		Str("service", "user").
		Str("module", "get user estamp").
		Str("user_id", req.UId).
		Msg("Successfully get user estamp")

	return &proto.GetUserEstampResponse{
		EventList: eventSrv.RawToDtoList(&events),
	}, nil
}

func DtoToRaw(in *proto.User) (result *user.User, err error) {
	var id uuid.UUID
	var groupId *uuid.UUID
	var baanId *uuid.UUID

	if in.Id != "" {
		id, err = uuid.Parse(in.Id)
		if err != nil {
			return nil, err
		}
	}

	if in.BaanId != "" {
		bId, err := uuid.Parse(in.BaanId)
		if err != nil {
			return nil, err
		}

		baanId = &bId

		if bId == uuid.Nil {
			baanId = nil
		}
	}

	return &user.User{
		Base: model.Base{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Title:           in.Title,
		Firstname:       in.Firstname,
		Lastname:        in.Lastname,
		Nickname:        in.Nickname,
		StudentID:       in.StudentID,
		Faculty:         in.Faculty,
		Year:            in.Year,
		Phone:           in.Phone,
		LineID:          in.LineID,
		Email:           in.Email,
		AllergyFood:     in.AllergyFood,
		FoodRestriction: in.FoodRestriction,
		AllergyMedicine: in.AllergyMedicine,
		Disease:         in.Disease,
		GroupID:         groupId,
		BaanID:          baanId,
		CanSelectBaan:   &in.CanSelectBaan,
	}, nil
}

func RawToDto(in *user.User, imgUrl string) *proto.User {
	var baanId string

	if in.IsVerify == nil {
		in.IsVerify = utils.BoolAdr(false)
	}

	if in.CanSelectBaan == nil {
		in.CanSelectBaan = utils.BoolAdr(false)
	}

	if in.IsGotTicket == nil {
		in.IsGotTicket = utils.BoolAdr(false)
	}

	if in.BaanID != nil {
		baanId = in.BaanID.String()
	}

	return &proto.User{
		Id:              in.ID.String(),
		Title:           in.Title,
		Firstname:       in.Firstname,
		Lastname:        in.Lastname,
		Nickname:        in.Nickname,
		StudentID:       in.StudentID,
		Faculty:         in.Faculty,
		Year:            in.Year,
		Phone:           in.Phone,
		LineID:          in.LineID,
		Email:           in.Email,
		AllergyFood:     in.AllergyFood,
		FoodRestriction: in.FoodRestriction,
		AllergyMedicine: in.AllergyMedicine,
		Disease:         in.Disease,
		ImageUrl:        imgUrl,
		CanSelectBaan:   *in.CanSelectBaan,
		IsVerify:        *in.IsVerify,
		IsGotTicket:     *in.IsGotTicket,
		BaanId:          baanId,
	}
}

func GetColumnName(verifyName string) string {
	switch verifyName {
	case "vaccine":
		return "is_verify"
	case "ticket":
		return "is_got_ticket"
	default:
		return ""
	}
}
