package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"
	"telegram-notification-api/internal/types/nulltypes"

	desc "telegram-notification-api/api"
)

func (s *server) GetUsersByFilter(
	ctx context.Context,
	req *desc.GetUsersByFilterRequest,
) (*desc.GetUsersByFilterResponse, error) {
	h, err := newGetUsersByFilterHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}

	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *getUsersByFilterHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}

	users, err := h.dao.NewUserQuery().
		GetUserByFilter(h.ctx, h.filterUser, h.limit, h.offset, h.fields...)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}

	h.filteredUsers = users
	return nil
}

func (h *getUsersByFilterHandler) response() *desc.GetUsersByFilterResponse {
	dest := make([]*desc.User, 0, len(h.filteredUsers))

	for _, user := range h.filteredUsers {
		var patronymic *string
		if user.Patronymic.Valid {
			patronymic = &user.Patronymic.String
		}

		dest = append(dest, &desc.User{
			UserId:                 user.Id,
			TelegramId:             user.TelegramId,
			UserRole:               desc.UserRole(desc.UserRole_value[user.Role]),
			UserNotificationStatus: desc.UserNotificationStatus(desc.UserNotificationStatus_value[user.NotificationStatus]),
			Group:                  user.Group,
			Fio: &desc.FIO{
				Firstname:  user.Firstname,
				Surname:    user.Surname,
				Patronymic: patronymic,
			},
			MobilePhone: user.MobilePhone,
			UserStatus:  desc.UserStatus(desc.UserNotificationStatus_value[user.Status]),
		})
	}
	return &desc.GetUsersByFilterResponse{
		Users:  dest,
		Limit:  int64(h.limit),
		Offset: int64(h.offset),
		Count:  int64(len(dest)),
	}
}

type getUsersByFilterHandler struct {
	ctx context.Context
	dao dao.DAO

	filterUser dao.UserTable
	fields     []string
	limit      uint64
	offset     uint64

	filteredUsers []dao.UserTable
}

func newGetUsersByFilterHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.GetUsersByFilterRequest,
) (*getUsersByFilterHandler, error) {
	h := getUsersByFilterHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *getUsersByFilterHandler) validate() error {
	if h.limit <= 0 {
		return errors.
			NewNetworkError(codes.InvalidArgument, "limit must be specified").
			ToGRPCError()
	}
	return nil
}

func (h *getUsersByFilterHandler) adapt(req *desc.GetUsersByFilterRequest) *getUsersByFilterHandler {
	if req.TelegramId != nil {
		h.filterUser.TelegramId = req.GetTelegramId()
		h.fields = append(h.fields, "telegram_id")
	}
	if req.Firstname != nil {
		h.filterUser.Firstname = req.GetFirstname()
		h.fields = append(h.fields, "firstname")
	}
	if req.Surname != nil {
		h.filterUser.Surname = req.GetSurname()
		h.fields = append(h.fields, "surname")
	}
	if req.Patronymic != nil {
		h.filterUser.Patronymic = nulltypes.NewNullString(req.Patronymic)
		h.fields = append(h.fields, "patronymic")
	}
	if req.MobilePhone != nil {
		h.filterUser.MobilePhone = req.GetMobilePhone()
		h.fields = append(h.fields, "mobile_phone")
	}
	if req.Group != nil {
		h.filterUser.Group = req.GetGroup()
		h.fields = append(h.fields, "user_group")
	}
	if req.UserNotificationStatus != nil {
		h.filterUser.NotificationStatus = req.GetUserNotificationStatus().String()
		h.fields = append(h.fields, "notification_status")
	}
	if req.UserStatus != nil {
		h.filterUser.Status = req.GetUserStatus().String()
		h.fields = append(h.fields, "status")
	}
	if req.UserRole != nil {
		h.filterUser.Role = req.GetUserRole().String()
		h.fields = append(h.fields, "role")
	}
	h.limit = uint64(req.GetLimit())
	h.offset = uint64(req.GetOffset())
	return h
}
