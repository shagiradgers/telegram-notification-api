package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"

	desc "telegram-notification-api/api"
)

func (s *server) GetNotifications(
	ctx context.Context,
	req *desc.GetNotificationsRequest,
) (*desc.GetNotificationsResponse, error) {
	h, err := newGetNotificationsHandler(ctx, s.dao, req)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *getNotificationsHandler) handle() error {
	if h == nil {
		return fmt.Errorf("got nil receiver")
	}

	notifications, err := h.dao.
		NewNotificationQuery().
		GetNotifications(h.ctx, h.notificationIDs, h.limit, h.offset)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	h.notifications = notifications
	return nil
}

func (h *getNotificationsHandler) response() *desc.GetNotificationsResponse {
	notifications := make([]*desc.Notification, 0, len(h.notifications))
	for idx := range h.notifications {
		var mediaContent *string
		if h.notifications[idx].MediaContent.Valid {
			mediaContent = &h.notifications[idx].MediaContent.String
		}

		notifications = append(notifications, &desc.Notification{
			NotificationId:     h.notifications[idx].ID,
			SenderId:           h.notifications[idx].SenderID,
			ReceiverIds:        h.notifications[idx].ReceiverIDs,
			Message:            h.notifications[idx].Message,
			MediaContent:       mediaContent,
			NotificationStatus: desc.NotificationStatus(desc.NotificationStatus_value[h.notifications[idx].Status]),
			Date:               timestamppb.New(h.notifications[idx].Date),
		})
	}

	return &desc.GetNotificationsResponse{
		Notification: notifications,
		Limit:        int64(h.limit),
		Offset:       int64(h.offset),
		Count:        int64(len(h.notifications)),
	}
}

type getNotificationsHandler struct {
	ctx context.Context
	dao dao.DAO

	notificationIDs []int64
	limit           uint64
	offset          uint64

	notifications []dao.NotificationTable
}

func newGetNotificationsHandler(
	ctx context.Context,
	dao dao.DAO,
	req *desc.GetNotificationsRequest,
) (*getNotificationsHandler, error) {
	h := &getNotificationsHandler{
		ctx: ctx,
		dao: dao,
	}
	return h.adapt(req), h.validate()
}

func (h *getNotificationsHandler) adapt(req *desc.GetNotificationsRequest) *getNotificationsHandler {
	h.offset = uint64(req.GetOffset())
	h.limit = uint64(req.GetLimit())
	h.notificationIDs = req.GetNotificationIds()
	return h
}

func (h *getNotificationsHandler) validate() error {
	if len(h.notificationIDs) == 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "notification_ids must be specified").
			ToGRPCError()
	}
	if h.limit <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "limit must be specified").
			ToGRPCError()
	}
	return nil
}
