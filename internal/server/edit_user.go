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

func (s *server) EditUser(
	ctx context.Context,
	req *desc.EditUserRequest,
) (*desc.EditUserResponse, error) {
	h, err := newEditUserHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}

	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *editUserHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}

	user, err := h.dao.NewUserQuery().ChangeUser(h.ctx, h.userToEdit, h.fields...)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	h.editedUser = user
	return nil
}

func (h *editUserHandler) response() *desc.EditUserResponse {
	var patronymic *string
	if h.editedUser.Patronymic.Valid {
		patronymic = &h.editedUser.Patronymic.String
	}

	return &desc.EditUserResponse{
		User: &desc.User{
			UserId:                 h.editedUser.Id,
			TelegramId:             h.editedUser.TelegramId,
			UserRole:               desc.UserRole(desc.UserRole_value[h.editedUser.Role]),
			UserNotificationStatus: desc.UserNotificationStatus(desc.UserNotificationStatus_value[h.editedUser.NotificationStatus]),
			Group:                  h.editedUser.Group,
			Fio: &desc.FIO{
				Firstname:  h.editedUser.Firstname,
				Surname:    h.editedUser.Surname,
				Patronymic: patronymic,
			},
			MobilePhone: h.editedUser.MobilePhone,
			UserStatus:  desc.UserStatus(desc.UserNotificationStatus_value[h.editedUser.Status]),
		},
	}
}

type editUserHandler struct {
	ctx        context.Context
	dao        dao.DAO
	fields     []string
	userToEdit dao.UserTable

	editedUser dao.UserTable
}

func newEditUserHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.EditUserRequest,
) (*editUserHandler, error) {
	h := editUserHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *editUserHandler) adapt(req *desc.EditUserRequest) *editUserHandler {
	h.userToEdit.Id = req.GetUserId()

	if req.UserStatus != nil {
		h.userToEdit.Status = req.GetUserStatus().String()
		h.fields = append(h.fields, "status")
	}
	if req.UserRole != nil {
		h.userToEdit.Role = req.GetUserRole().String()
		h.fields = append(h.fields, "role")
	}
	if req.UserNotificationStatus != nil {
		h.userToEdit.NotificationStatus = req.GetUserNotificationStatus().String()
		h.fields = append(h.fields, "notification_status")
	}
	if req.TelegramId != nil {
		h.userToEdit.TelegramId = req.GetTelegramId()
		h.fields = append(h.fields, "telegram_id")
	}
	if req.Group != nil {
		h.userToEdit.Group = req.GetGroup()
		h.fields = append(h.fields, "user_group")
	}
	if req.Firstname != nil {
		h.userToEdit.Firstname = req.GetFirstname()
		h.fields = append(h.fields, "firstname")
	}
	if req.Surname != nil {
		h.userToEdit.Surname = req.GetSurname()
		h.fields = append(h.fields, "surname")
	}
	if req.Patronymic != nil {
		h.userToEdit.Patronymic = nulltypes.NewNullString(req.Patronymic)
		h.fields = append(h.fields, "patronymic")
	}
	if req.MobilePhone != nil {
		h.userToEdit.MobilePhone = req.GetMobilePhone()
		h.fields = append(h.fields, "mobile_phone")
	}

	return h
}

func (h *editUserHandler) validate() error {
	if h.userToEdit.Id <= 0 {
		return errors.
			NewNetworkError(codes.InvalidArgument, "telegram_id must be specified").
			ToGRPCError()
	}
	if h.userToEdit.TelegramId != 0 && h.userToEdit.TelegramId < 0 {
		return errors.
			NewNetworkError(codes.InvalidArgument, "telegram_id must be specified").
			ToGRPCError()
	}

	return nil
}
