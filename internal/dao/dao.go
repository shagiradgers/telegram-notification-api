package dao

import (
	sq "github.com/Masterminds/squirrel"
	"telegram-notification-api/internal/storage"
)

type DAO interface {
	NewNotificationQuery() NotificationQuery
	NewUserQuery() UserQuery

	Close() error
}

type dao struct {
	db storage.Storage
}

func (d *dao) NewNotificationQuery() NotificationQuery {
	return newNotificationQuery(d.db)
}

func (d *dao) NewUserQuery() UserQuery {
	return newUserQuery(d.db)
}

func (d *dao) Close() error {
	if d.db == nil {
		return nil
	}
	return d.db.Close()
}

func NewDAO(s storage.Storage) DAO {
	return &dao{
		db: s,
	}
}

func qb() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
