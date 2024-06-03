package server

import (
	"context"

	desc "telegram-notification-api/api"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"
)

func (s *server) GetGroups(
	ctx context.Context,
	_ *desc.GetGroupsRequest,
) (*desc.GetGroupsResponse, error) {
	h, err := newGetGroupsHandler(ctx, s.dao)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	return h.response(), err
}

func (h *getGroupsHandler) handle() error {
	var err error
	h.groups, err = h.dao.NewUserQuery().GetGroups(h.ctx)
	return errors.WrapToNetwork(err).ToGRPCError()
}

func (h *getGroupsHandler) response() *desc.GetGroupsResponse {
	return &desc.GetGroupsResponse{
		Groups: h.groups,
		Count:  int64(len(h.groups)),
	}
}

type getGroupsHandler struct {
	ctx context.Context
	dao dao.DAO

	groups []string
}

func newGetGroupsHandler(
	ctx context.Context,
	dao dao.DAO,
) (*getGroupsHandler, error) {
	h := &getGroupsHandler{
		ctx: ctx,
		dao: dao,
	}
	return h, nil
}
