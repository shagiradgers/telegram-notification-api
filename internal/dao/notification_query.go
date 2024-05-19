package dao

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	desc "telegram-notification-api/api"
	"telegram-notification-api/internal/storage"
	"time"
)

type NotificationQuery interface {
	GetNotification(ctx context.Context, ID int64) (NotificationTable, error)
	CreateNotification(
		ctx context.Context,
		senderID int64,
		receiverIDs []int64,
		message string,
		mediaContent sql.NullString,
		date time.Time,
	) (NotificationTable, error)
	GetNotifications(
		ctx context.Context,
		IDs []int64,
		limit uint64,
		offset uint64,
	) ([]NotificationTable, error)
	UpdateNotificationStatus(
		ctx context.Context,
		notificationID int64,
		status string,
	) error
}

type notificationQuery struct {
	db storage.Storage
}

func newNotificationQuery(db storage.Storage) NotificationQuery {
	return &notificationQuery{
		db: db,
	}
}

func (n *notificationQuery) GetNotification(
	ctx context.Context,
	ID int64,
) (NotificationTable, error) {
	var dest NotificationTable
	query := qb().
		Select(dest.columns()...).
		From(notificationTableName).
		Where(sq.Eq{"id": ID})

	err := n.db.GetX(ctx, &dest, query)
	return dest, err
}

func (n *notificationQuery) CreateNotification(
	ctx context.Context,
	senderID int64,
	receiverIDs []int64,
	message string,
	mediaContent sql.NullString,
	date time.Time,
) (NotificationTable, error) {
	var dest NotificationTable

	query := qb().
		Insert(notificationTableName).
		Columns(
			"sender_id",
			"receiver_ids",
			"message",
			"media_content",
			"date",
			"status",
		).
		Values(
			senderID,
			pq.Array(receiverIDs),
			message,
			mediaContent,
			date,
			desc.NotificationStatus_CREATED.String(),
		).
		Suffix("RETURNING *")

	err := n.db.GetX(ctx, &dest, query)
	return dest, err
}

func (n *notificationQuery) GetNotifications(
	ctx context.Context,
	IDs []int64,
	limit uint64,
	offset uint64,
) ([]NotificationTable, error) {
	var dest []NotificationTable
	query := qb().
		Select(NotificationTable{}.columns()...).
		From(notificationTableName).
		Where("id = ANY(?)", pq.Array(IDs)).
		Offset(offset).
		Limit(limit)

	err := n.db.SelectX(ctx, &dest, query)
	return dest, err
}

func (n *notificationQuery) UpdateNotificationStatus(
	ctx context.Context,
	notificationID int64,
	status string,
) error {
	query := qb().
		Update(notificationTableName).
		Set("status", status).
		Where(sq.Eq{"id": notificationID})
	return n.db.ExecX(ctx, query)
}
