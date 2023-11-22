package store

import (
	"context"
	"fmt"

	"github.com/l-orlov/simple-todo-list/server/model"
)

// todo: add timeout for ctx

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
