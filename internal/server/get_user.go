package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"

	desc "telegram-notification-api/api"
)

func (s *server) GetUser(
	ctx context.Context,
	req *desc.GetUserRequest,
) (*desc.GetUserResponse, error) {
	h, err := newGetUserHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *getUserHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}
	u, err := h.dao.NewUserQuery().GetUser(h.ctx, h.userID)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}

	h.Id = u.Id
	h.TelegramId = u.TelegramId
	h.Role = desc.UserRole_value[u.Role]
	h.NotificationStatus = desc.UserNotificationStatus_value[u.NotificationStatus]
	h.Group = u.Group
	h.Firstname = u.Firstname
	h.Surname = u.Surname
	if u.Patronymic.Valid {
		patronymic := u.Patronymic.String
		h.Patronymic = &patronymic
	}
	h.MobilePhone = u.MobilePhone
	h.Status = desc.UserStatus_value[u.Status]

	return nil
}

func (h *getUserHandler) response() *desc.GetUserResponse {
	return &desc.GetUserResponse{
		User: &desc.User{
			UserId:                 h.Id,
			TelegramId:             h.TelegramId,
			UserRole:               desc.UserRole(h.Role),
			UserNotificationStatus: desc.UserNotificationStatus(h.NotificationStatus),
			Group:                  h.Group,
			Fio: &desc.FIO{
				Firstname:  h.Firstname,
				Surname:    h.Surname,
				Patronymic: h.Patronymic,
			},
			MobilePhone: h.MobilePhone,
			UserStatus:  desc.UserStatus(h.Status),
		},
	}
}

type getUserHandler struct {
	ctx context.Context
	dao dao.DAO

	userID int64

	Id                 int64
	TelegramId         int64
	Role               int32
	NotificationStatus int32
	Group              string
	Firstname          string
	Surname            string
	Patronymic         *string
	MobilePhone        string
	Status             int32
}

func newGetUserHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.GetUserRequest,
) (*getUserHandler, error) {
	h := getUserHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *getUserHandler) adapt(req *desc.GetUserRequest) *getUserHandler {
	h.userID = req.GetUserId()
	return h
}

func (h *getUserHandler) validate() error {
	if h.userID <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "user_id must be specified").
			ToGRPCError()
	}
	return nil
}
