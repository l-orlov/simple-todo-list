package store

import (
	"context"
	"fmt"

	"github.com/l-orlov/simple-todo-list/server/model"
)

// todo: add timeout for ctx

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
		return fmt.Errorf("db.QueryRowContext and Scan: %w", err)
	}

	return nil
}

func (s *Storage) GetTasks(ctx context.Context) ([]*model.Task, error) {
	queryBuilder := psql().
		Select(asteriskTasks).
		From(model.Task{}.DbTable())

	// Получаем SQL-запрос и его аргументы
	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("queryBuilder.ToSql: %w", err)
	}

	// Выполняем запрос
	rows, err := s.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	defer rows.Close()

	// Обрабатываем результат запроса
	var tasks []*model.Task
	for rows.Next() {
		task := &model.Task{}
		err = rows.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		tasks = append(tasks, task)
	}

	// Проверяем наличие ошибок после выполнения запроса
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return tasks, nil
}

func taskAttrs(record *model.Task) map[string]interface{} {
	return map[string]interface{}{
		"title":  record.Title,
		"status": record.Status,
	}
}
