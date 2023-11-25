package store

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/l-orlov/simple-todo-list/server/internal/model"
)

func (s *Storage) CreateUser(ctx context.Context, record *model.User) error {
	// Создаем запрос вставки
	queryBuilder := psql().
		Insert(record.DbTable()).
		SetMap(userAttrs(record)).
		Suffix("RETURNING id, created_at, updated_at")

	// Получаем SQL-запрос и его аргументы
	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("queryBuilder.ToSql: %w", err)
	}

	// Выполняем запрос и получаем ID вставленной записи
	err = s.db.QueryRowContext(ctx, sqlQuery, args...).Scan(&record.ID, &record.CreatedAt, &record.UpdatedAt)
	if err != nil {
		return dbError(err)
	}

	return nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	record := &model.User{}

	// Создаем запрос вставки
	queryBuilder := psql().
		Select(asteriskUsers).
		From(record.DbTable()).
		Where(sq.Eq{"email": email}).
		Limit(1)

	// Получаем SQL-запрос и его аргументы
	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("queryBuilder.ToSql: %w", err)
	}

	// Выполняем запрос и получаем ID вставленной записи
	err = s.db.QueryRowContext(ctx, sqlQuery, args...).Scan(&record.ID, &record.Email, &record.Password, &record.CreatedAt, &record.UpdatedAt)
	if err != nil {
		return nil, dbError(err)
	}

	return record, nil
}

func userAttrs(record *model.User) map[string]interface{} {
	return map[string]interface{}{
		"email":      record.Email,
		"password":   record.Password,
		"updated_at": "now()",
	}
}
