package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"

	desc "telegram-notification-api/api"
)

func (s *server) DeleteUser(
	ctx context.Context,
	req *desc.DeleteUserRequest,
) (*desc.DeleteUserResponse, error) {
	h, err := newDeleteUserHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *deleteUserHandler) response() *desc.DeleteUserResponse {
	return &desc.DeleteUserResponse{}
}

func (h *deleteUserHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}
	err := h.dao.NewUserQuery().DeleteUser(h.ctx, h.userID)
	return errors.WrapToNetwork(err).ToGRPCError()
}

type deleteUserHandler struct {
	ctx context.Context
	dao dao.DAO

	userID int64
}

func newDeleteUserHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.DeleteUserRequest,
) (deleteUserHandler, error) {
	h := deleteUserHandler{
		ctx:    ctx,
		dao:    dao,
		userID: req.GetUserId(),
	}
	return h, h.validate()
}

func (h *deleteUserHandler) validate() error {
	if h.userID <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "user_id must be specified").
			ToGRPCError()
	}
	return nil
}
