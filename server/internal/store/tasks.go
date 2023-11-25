package store

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/l-orlov/simple-todo-list/server/internal/model"
)

func (s *Storage) CreateTask(ctx context.Context, record *model.Task) error {
	// Создаем запрос вставки
	queryBuilder := psql().
		Insert(record.DbTable()).
		SetMap(taskAttrs(record)).
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

func (s *Storage) UpdateTaskByID(ctx context.Context, record *model.Task) error {
	// Создаем запрос вставки
	queryBuilder := psql().
		Update(record.DbTable()).
		SetMap(taskAttrs(record)).
		Where(sq.Eq{
			"id":      record.ID,
			"user_id": record.UserID,
		}).
		Suffix("RETURNING " + asteriskTasks)

	// Получаем SQL-запрос и его аргументы
	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("queryBuilder.ToSql: %w", err)
	}

	// Выполняем запрос и получаем ID вставленной записи
	err = s.db.QueryRowContext(ctx, sqlQuery, args...).Scan(&record.ID, &record.UserID, &record.Title, &record.Status, &record.CreatedAt, &record.UpdatedAt)
	if err != nil {
		return dbError(err)
	}

	return nil
}

// GetTasksByUserID делает поиск тасок по user_id
func (s *Storage) GetTasksByUserID(ctx context.Context, userID uuid.UUID) ([]*model.Task, error) {
	queryBuilder := psql().
		Select(asteriskTasks).
		From(model.Task{}.DbTable()).
		Where(sq.Eq{"user_id": userID}).
		// Получаем все таски кроме удаленных
		Where(sq.NotEq{"status": model.TaskStatusDeleted}) // status <> 4

	// Получаем SQL-запрос и его аргументы
	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("queryBuilder.ToSql: %w", err)
	}

	// Выполняем запрос
	rows, err := s.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, dbError(err)
	}
	defer rows.Close()

	// Обрабатываем результат запроса
	records := make([]*model.Task, 0)
	for rows.Next() {
		record := &model.Task{}
		err = rows.Scan(&record.ID, &record.UserID, &record.Title, &record.Status, &record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		records = append(records, record)
	}

	// Проверяем наличие ошибок после выполнения запроса
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return records, nil
}

func taskAttrs(record *model.Task) map[string]interface{} {
	return map[string]interface{}{
		"user_id":    record.UserID,
		"title":      record.Title,
		"status":     record.Status,
		"updated_at": "now()",
	}
}
