package server

import (
	"context"
	"database/sql"
	"fmt"
	"telegram-notification-api/internal/types/nulltypes"

	"google.golang.org/grpc/codes"
	desc "telegram-notification-api/api"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"
)

func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	h, err := newCreateUserHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	if err != nil {
		return nil, err
	}
	return &desc.CreateUserResponse{}, nil
}

func (h *createUserHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}
	err := h.dao.
		NewUserQuery().
		CreateUser(
			h.ctx,
			h.TelegramId,
			h.UserRole,
			h.UserNotificationStatus,
			h.Group,
			h.Fio.Firstname,
			h.Fio.Surname,
			h.Fio.Patronymic,
			h.MobilePhone,
			h.UserStatus,
		)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	return nil
}

type createUserHandler struct {
	ctx context.Context
	dao dao.DAO

	TelegramId             int64
	UserRole               string
	UserNotificationStatus string
	Group                  string
	Fio                    createUserHandlerFIO
	MobilePhone            string
	UserStatus             string
}

type createUserHandlerFIO struct {
	Firstname  string
	Surname    string
	Patronymic sql.NullString
}

func newCreateUserHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.CreateUserRequest,
) (*createUserHandler, error) {
	h := createUserHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *createUserHandler) adapt(req *desc.CreateUserRequest) *createUserHandler {
	h.TelegramId = req.GetTelegramId()
	h.UserRole = req.GetUserRole().String()
	h.UserNotificationStatus = req.GetUserNotificationStatus().String()
	h.Group = req.GetGroup()
	h.MobilePhone = req.GetMobilePhone()
	h.UserStatus = desc.UserStatus_ACTIVE.String()
	h.Fio = createUserHandlerFIO{
		Firstname:  req.GetFio().GetFirstname(),
		Surname:    req.GetFio().GetSurname(),
		Patronymic: nulltypes.NewNullString(req.GetFio().Patronymic),
	}
	return h
}

func (h *createUserHandler) validate() error {
	if h.TelegramId <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "telegram_id must be specified").
			ToGRPCError()
	}
	if h.Group == "" {
		return errors.NewNetworkError(codes.InvalidArgument, "group must be specified").
			ToGRPCError()
	}
	if h.MobilePhone == "" {
		return errors.NewNetworkError(codes.InvalidArgument, "mobile_phone must be specified").
			ToGRPCError()
	}
	if h.Fio.Firstname == "" {
		return errors.NewNetworkError(codes.InvalidArgument, "fio.firstname must be specified").
			ToGRPCError()
	}
	if h.Fio.Surname == "" {
		return errors.NewNetworkError(codes.InvalidArgument, "fio.surname must be specified").
			ToGRPCError()
	}
	return nil
}
