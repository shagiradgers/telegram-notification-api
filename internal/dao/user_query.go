package dao

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"telegram-notification-api/internal/storage"
)

type UserQuery interface {
	GetUser(ctx context.Context, userID int64) (UserTable, error)
	GetUsersByIds(ctx context.Context, userIDs []int64, limit uint64, offset uint64) ([]UserTable, error)
	CreateUser(
		ctx context.Context,
		telegramID int64,
		userRole string,
		userNotificationStatus string,
		group string,
		firstname string,
		lastname string,
		patronymic sql.NullString,
		mobilePhone string,
		userStatus string,
	) error
	DeleteUser(ctx context.Context, userID int64) error
	ChangeUser(ctx context.Context, user UserTable, fields ...string) (UserTable, error)
	GetUserByFilter(
		ctx context.Context,
		user UserTable,
		limit uint64,
		offset uint64,
		fields ...string,
	) ([]UserTable, error)
	GetGroups(ctx context.Context) ([]string, error)
}

type userQuery struct {
	storage storage.Storage
}

func newUserQuery(storage storage.Storage) UserQuery {
	return &userQuery{storage: storage}
}

func (u *userQuery) GetUser(ctx context.Context, userID int64) (UserTable, error) {
	var dest UserTable
	query := qb().
		Select(dest.columns()...).
		From(userTableName).
		Where(sq.Eq{"id": userID})

	err := u.storage.GetX(ctx, &dest, query)
	return dest, err
}

func (u *userQuery) CreateUser(
	ctx context.Context,
	telegramID int64,
	role string,
	notificationStatus string,
	group string,
	firstname string,
	lastname string,
	patronymic sql.NullString,
	mobilePhone string,
	userStatus string,
) error {
	query := qb().
		Insert(userTableName).
		Columns(
			"telegram_id",
			"role",
			"notification_status",
			"user_group",
			"firstname",
			"surname",
			"patronymic",
			"mobile_phone",
			"status",
		).
		Values(
			telegramID,
			role,
			notificationStatus,
			group,
			firstname,
			lastname,
			patronymic,
			mobilePhone,
			userStatus,
		)
	err := u.storage.ExecX(ctx, query)
	return err
}

func (u *userQuery) DeleteUser(ctx context.Context, userID int64) error {
	query := qb().
		Delete(userTableName).
		Where(sq.Eq{"id": userID})
	return u.storage.ExecX(ctx, query)
}

func (u *userQuery) GetUsersByIds(ctx context.Context, userIDs []int64, limit uint64, offset uint64) ([]UserTable, error) {
	var dest []UserTable
	query := qb().
		Select(UserTable{}.columns()...).
		From(userTableName).
		Where("id = ANY(?)", pq.Array(userIDs)).
		Limit(limit).
		Offset(offset)
	err := u.storage.SelectX(ctx, &dest, query)
	if err != nil {
		return nil, err
	}

	return dest, err
}

func (u *userQuery) ChangeUser(ctx context.Context, user UserTable, fields ...string) (UserTable, error) {
	var dest UserTable
	userMap := user.toMap()

	query := qb().
		Update(userTableName).
		Where(sq.Eq{"id": user.Id})

	for _, field := range fields {
		query = query.Set(field, userMap[field])
	}
	query = query.Suffix("RETURNING *")

	err := u.storage.GetX(ctx, &dest, query)
	return dest, err
}

func (u *userQuery) GetUserByFilter(
	ctx context.Context,
	user UserTable,
	limit uint64,
	offset uint64,
	fields ...string,
) ([]UserTable, error) {
	var dest []UserTable
	userMap := user.toMap()

	query := qb().
		Select(UserTable{}.columns()...).
		From(userTableName)

	for _, field := range fields {
		query = query.Where(sq.Eq{field: userMap[field]})
	}
	query = query.
		Limit(limit).
		Offset(offset)

	err := u.storage.SelectX(ctx, &dest, query)
	return dest, err
}

func (q *userQuery) GetGroups(ctx context.Context) ([]string, error) {
	var dest []string

	query := qb().
		Select("DISTINCT user_group").
		From(userTableName)

	err := q.storage.SelectX(ctx, &dest, query)
	return dest, err
}
