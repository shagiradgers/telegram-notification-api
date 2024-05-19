package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage interface {
	Exec(sq squirrel.Sqlizer) error
	ExecX(ctx context.Context, sq squirrel.Sqlizer) error
	Get(dest interface{}, sq squirrel.Sqlizer) error
	GetX(ctx context.Context, dest interface{}, sq squirrel.Sqlizer) error
	Select(dest interface{}, sq squirrel.Sqlizer) error
	SelectX(ctx context.Context, dest interface{}, sq squirrel.Sqlizer) error
	Close() error
}

type storage struct {
	db *sqlx.DB
}

func NewStorage(dataSourceName string) (Storage, error) {
	const driverName = "postgres"
	db, err := sqlx.Open(driverName, dataSourceName)
	return &storage{
		db: db,
	}, err
}

func (s *storage) ExecX(ctx context.Context, sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, sql, args...)
	return err
}

func (s *storage) Exec(sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	_, err = s.db.Exec(sql, args...)
	return err
}

func (s *storage) SelectX(ctx context.Context, dest interface{}, sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	return s.db.SelectContext(ctx, dest, sql, args...)
}

func (s *storage) Select(dest interface{}, sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	return s.db.Select(dest, sql, args...)
}

func (s *storage) GetX(ctx context.Context, dest interface{}, sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	return s.db.GetContext(ctx, dest, sql, args...)
}

func (s *storage) Get(dest interface{}, sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return err
	}
	return s.db.Get(dest, sql, args...)
}

func (s *storage) Close() error {
	return s.db.Close()
}
