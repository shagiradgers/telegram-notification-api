package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"telegram-notification-api/internal/clients"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/errors"
	"telegram-notification-api/internal/types/nulltypes"
	"time"

	desc "telegram-notification-api/api"
)

func (s *server) SendNotification(
	ctx context.Context,
	req *desc.SendNotificationRequest,
) (*desc.SendNotificationResponse, error) {
	h, err := newSendNotificationHandler(ctx, s.dao, s.clients, req)
	if err != nil {
		return nil, err
	}
	err = h.handle()
	if err != nil {
		return nil, err
	}
	return h.response(), nil
}

func (h *sendNotificationHandler) handle() error {
	if h == nil {
		return fmt.Errorf("go nil receiver")
	}
	notification, err := h.dao.
		NewNotificationQuery().
		CreateNotification(
			h.ctx,
			h.senderID,
			h.receiverIDs,
			h.message,
			nulltypes.NewNullString(h.mediaContent),
			h.now,
		)
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	h.createdNotification = notification
	if err = h.sendMessageToTelegram(); err != nil {
		// если мы не смогли отправить сообщение, надо проставить проблему
		if err := h.dao.
			NewNotificationQuery().
			UpdateNotificationStatus(
				h.ctx,
				h.createdNotification.ID,
				desc.NotificationStatus_PROBLEM.String(),
			); err != nil {
			return errors.WrapToNetwork(err).ToGRPCError()
		}
		return errors.WrapToNetwork(err).ToGRPCError()
	}

	err = h.dao.
		NewNotificationQuery().
		UpdateNotificationStatus(h.ctx, h.createdNotification.ID, desc.NotificationStatus_SEND.String())
	if err != nil {
		return errors.WrapToNetwork(err).ToGRPCError()
	}
	h.createdNotification.Status = desc.NotificationStatus_SEND.String()
	return nil
}

func (h *sendNotificationHandler) sendMessageToTelegram() error {
	var err error
	var user dao.UserTable

	for idx := range h.receiverIDs {
		user, err = h.dao.NewUserQuery().GetUser(h.ctx, h.receiverIDs[idx])
		if err != nil {
			return err
		}
		err = h.clients.TelegramClient().SendMessage(
			h.ctx,
			user.TelegramId,
			h.message,
			h.mediaContent,
			!desc.UserNotificationStatus(desc.UserNotificationStatus_value[user.NotificationStatus]).
				ToBool(),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *sendNotificationHandler) response() *desc.SendNotificationResponse {
	return &desc.SendNotificationResponse{
		NotificationId: h.createdNotification.ID,
		MessageStatus:  desc.NotificationStatus(desc.NotificationStatus_value[h.createdNotification.Status]),
	}
}

type sendNotificationHandler struct {
	ctx     context.Context
	dao     dao.DAO
	clients clients.Clients

	senderID     int64
	receiverIDs  []int64
	message      string
	now          time.Time
	mediaContent *string

	createdNotification dao.NotificationTable
}

func newSendNotificationHandler(
	ctx context.Context,
	dao dao.DAO,
	clients clients.Clients,
	req *desc.SendNotificationRequest,
) (*sendNotificationHandler, error) {
	h := &sendNotificationHandler{
		ctx:     ctx,
		dao:     dao,
		clients: clients,
		now:     time.Now(),
	}
	return h.adapt(req), h.validate()
}

func (h *sendNotificationHandler) validate() error {
	if h.senderID <= 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "sender_id must be specified").
			ToGRPCError()
	}
	if len(h.receiverIDs) == 0 {
		return errors.NewNetworkError(codes.InvalidArgument, "receiver_ids must be specified").
			ToGRPCError()
	}
	if h.message == "" {
		return errors.NewNetworkError(codes.InvalidArgument, "message must be specified").
			ToGRPCError()
	}
	if h.mediaContent != nil && *h.mediaContent == "" {
		return errors.NewNetworkError(codes.InvalidArgument, "media_content must be specified").
			ToGRPCError()
	}
	return nil
}

func (h *sendNotificationHandler) adapt(req *desc.SendNotificationRequest) *sendNotificationHandler {
	h.mediaContent = req.MediaContent
	h.senderID = req.GetSenderId()
	h.receiverIDs = req.GetReceiverIds()
	h.message = req.GetMessage()
	return h
}
