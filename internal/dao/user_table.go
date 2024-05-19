package dao

import (
	"database/sql"
	"github.com/elgris/stom"
)

const (
	userTableName = "users"
)

type UserTable struct {
	Id                 int64          `db:"id"`
	TelegramId         int64          `db:"telegram_id"`
	Role               string         `db:"role"`
	NotificationStatus string         `db:"notification_status"`
	Group              string         `db:"user_group"`
	Firstname          string         `db:"firstname"`
	Surname            string         `db:"surname"`
	Patronymic         sql.NullString `db:"patronymic"`
	MobilePhone        string         `db:"mobile_phone"`
	Status             string         `db:"status"`
}

var userTableStom = stom.MustNewStom(UserTable{})

func (u UserTable) columns() []string {
	return userTableStom.TagValues()
}

func (u UserTable) toMap() map[string]interface{} {
	m, err := userTableStom.ToMap(u)
	if err != nil {
		panic(err)
	}
	return m
}
