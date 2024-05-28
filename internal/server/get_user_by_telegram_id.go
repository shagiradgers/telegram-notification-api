package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	desc "telegram-notification-api/api"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"
)

func (s *server) GetUserByTelegramID(
	ctx context.Context,
	req *desc.GetUserByTelegramIDRequest,
) (*desc.GetUserByTelegramIDResponse, error) {
	h, err := newGetUserByTelegramIDHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *getUserByTelegramIDHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}

	users, err := h.dao.NewUserQuery().GetUserByFilter(
		h.ctx,
		dao.UserTable{TelegramId: h.telegramID},
		1,
		0,
		"telegram_id",
	)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	if len(users) == 0 {
		return errors.NewNetworkError(codes.NotFound, "user not found").ToGRPCError()
	}

	h.gotUser = users[0]
	return nil
}

func (h *getUserByTelegramIDHandler) response() *desc.GetUserByTelegramIDResponse {
	var patronymic *string
	if h.gotUser.Patronymic.Valid {
		patronymic = &h.gotUser.Patronymic.String
	}

	return &desc.GetUserByTelegramIDResponse{
		User: &desc.User{
			UserId:                 h.gotUser.Id,
			TelegramId:             h.gotUser.TelegramId,
			UserRole:               desc.UserRole(desc.UserRole_value[h.gotUser.Role]),
			UserNotificationStatus: desc.UserNotificationStatus(desc.UserNotificationStatus_value[h.gotUser.NotificationStatus]),
			Group:                  h.gotUser.Group,
			Fio: &desc.FIO{
				Firstname:  h.gotUser.Firstname,
				Surname:    h.gotUser.Surname,
				Patronymic: patronymic,
			},
			MobilePhone: h.gotUser.MobilePhone,
			UserStatus:  desc.UserStatus(desc.UserNotificationStatus_value[h.gotUser.Status]),
		},
	}
}

type getUserByTelegramIDHandler struct {
	ctx context.Context
	dao dao.DAO

	telegramID int64

	gotUser dao.UserTable
}

func newGetUserByTelegramIDHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.GetUserByTelegramIDRequest,
) (*getUserByTelegramIDHandler, error) {
	h := &getUserByTelegramIDHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *getUserByTelegramIDHandler) adapt(req *desc.GetUserByTelegramIDRequest) *getUserByTelegramIDHandler {
	h.telegramID = req.GetTelegramId()
	return h
}

func (h *getUserByTelegramIDHandler) validate() error {
	if h.telegramID <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "telegram_id must be specified").
			ToGRPCError()
	}
	return nil
}
