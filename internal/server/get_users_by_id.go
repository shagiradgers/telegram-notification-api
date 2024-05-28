package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"

	desc "telegram-notification-api/api"
)

func (s *server) GetUsersById(
	ctx context.Context,
	req *desc.GetUsersByIdRequest,
) (*desc.GetUsersByIdResponse, error) {
	h, err := newGetUsersByIdHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}

	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *getUserByIdHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}

	users, err := h.dao.NewUserQuery().GetUsersByIds(h.ctx, h.userIDs, uint64(h.limit), uint64(h.offset))
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	h.users = users
	return nil
}

func (h *getUserByIdHandler) response() *desc.GetUsersByIdResponse {
	resp := &desc.GetUsersByIdResponse{
		Users:  make([]*desc.User, 0, len(h.users)),
		Limit:  h.limit,
		Offset: h.offset,
		Count:  int64(len(h.users)),
	}

	for idx := range h.users {
		resp.Users = append(resp.Users, &desc.User{
			UserId:                 h.users[idx].Id,
			TelegramId:             h.users[idx].TelegramId,
			UserRole:               desc.UserRole(desc.UserRole_value[h.users[idx].Role]),
			UserNotificationStatus: desc.UserNotificationStatus(desc.UserNotificationStatus_value[h.users[idx].NotificationStatus]),
			Group:                  h.users[idx].Group,
			Fio: &desc.FIO{
				Firstname: h.users[idx].Firstname,
				Surname:   h.users[idx].Surname,
			},
			MobilePhone: h.users[idx].MobilePhone,
			UserStatus:  desc.UserStatus(desc.UserStatus_value[h.users[idx].Status]),
		})
		if h.users[idx].Patronymic.Valid {
			resp.Users[idx].Fio.Patronymic = &h.users[idx].Patronymic.String
		}
	}
	return resp
}

type getUserByIdHandler struct {
	ctx context.Context
	dao dao.DAO

	userIDs []int64
	limit   int64
	offset  int64

	users []dao.UserTable
}

func newGetUsersByIdHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.GetUsersByIdRequest,
) (*getUserByIdHandler, error) {
	h := getUserByIdHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *getUserByIdHandler) validate() error {
	if len(h.userIDs) == 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "user_ids must be specified").ToGRPCError()
	}
	if h.limit <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "limit must be specified").ToGRPCError()
	}
	if h.offset < 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "offset must be specified").ToGRPCError()
	}
	return nil
}

func (h *getUserByIdHandler) adapt(req *desc.GetUsersByIdRequest) *getUserByIdHandler {
	h.userIDs = req.GetUserIds()
	h.limit = req.GetLimit()
	h.offset = req.GetOffset()
	return h
}
