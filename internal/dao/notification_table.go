package dao

import (
	"database/sql"
	"github.com/elgris/stom"
	"github.com/lib/pq"
	"time"
)

type NotificationTable struct {
	ID           int64          `db:"id"`
	SenderID     int64          `db:"sender_id"`
	ReceiverIDs  pq.Int64Array  `db:"receiver_ids"`
	Message      string         `db:"message"`
	MediaContent sql.NullString `db:"media_content"`
	Status       string         `db:"status"`
	Date         time.Time      `db:"date"`
}

const (
	notificationTableName = "notifications"
)

var notificationTableStom = stom.MustNewStom(NotificationTable{})

func (t NotificationTable) columns() []string {
	return notificationTableStom.TagValues()
}

func (t NotificationTable) toMap() map[string]interface{} {
	m, err := notificationTableStom.ToMap(t)
	if err != nil {
		panic(err)
	}
	return m
}
