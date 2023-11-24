package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/l-orlov/simple-todo-list/server/model"
)

// CreateTask создает новую таску
// @summary Создает нового пользователя
// @description Хэндлер принимает POST-запрос с данными новой таски, делает запись в БД и возвращает результат в формате JSON.
// @tags users
// @accept json
// @produce json
// @param task body Task true "Данные новой таски"
// @success 200 {object} map[string]interface{} "JSON-ответ"
// @failure 400 {object} map[string]interface{} "JSON-ответ с сообщением об ошибке"
// @failure 500 {object} map[string]interface{} "JSON-ответ с сообщением об ошибке"
// @router /createUser [post]
func (c *Controller) CreateTask(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.CreateTask"

	// Проверяем метод запроса (должен быть POST)
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON-данные из тела запроса
	task := &model.Task{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(task); err != nil {
		http.Error(w, "invalid json in body", http.StatusBadRequest)
		return
	}

	// todo: add timeout
	ctx := r.Context()

	err := c.storage.CreateTask(ctx, task)
	if err != nil {
		log.Printf("%s: storage.CreateTask: %s", msgPrefix, err)
		http.Error(w, "error creating task", http.StatusInternalServerError)
		return
	}

	// Кодируем JSON и отправляем в ответе
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&task); err != nil {
		log.Printf("%s: json encode response: %s", msgPrefix, err)
		http.Error(w, "error encoding json response", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) UpdateTask(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.CreateTask"

	// Проверяем метод запроса (должен быть PUT)
	if r.Method != http.MethodPut {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON-данные из тела запроса
	task := &model.Task{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(task); err != nil {
		http.Error(w, "invalid json in body", http.StatusBadRequest)
		return
	}

	// todo: add timeout
	ctx := r.Context()

	err := c.storage.UpdateTaskByID(ctx, task)
	if err != nil {
		log.Printf("%s: storage.UpdateTaskByID: %s", msgPrefix, err)
		http.Error(w, "error updating task", http.StatusInternalServerError)
		return
	}

	// Кодируем JSON и отправляем в ответе
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&task); err != nil {
		log.Printf("%s: json encode response: %s", msgPrefix, err)
		http.Error(w, "error encoding json response", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetTasks(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.GetTasks"

	// Проверяем метод запроса (должен быть GET)
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Устанавливаем статус 200
	w.WriteHeader(http.StatusOK)

	// todo: add timeout
	ctx := r.Context()

	tasks, err := c.storage.GetTasks(ctx)
	if err != nil {
		log.Printf("%s: storage.GetTasks: %s", msgPrefix, err)
		http.Error(w, "error getting task", http.StatusInternalServerError)
		return
	}

	// Кодируем JSON и отправляем в ответе
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&tasks); err != nil {
		log.Printf("%s: json encode response: %s", msgPrefix, err)
		http.Error(w, "error encoding json response", http.StatusInternalServerError)
		return
	}
}
