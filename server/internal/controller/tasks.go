package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/l-orlov/simple-todo-list/server/internal/jwttoken"
	"github.com/l-orlov/simple-todo-list/server/internal/model"
)

// CreateTask создает новую таску
// @Summary Создание таски
// @Description Создает новую таску. В запросе нужно передать Bearer Token
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body model.TaskToCreate true "Данные новой таски"
// @Success 200 {string} string "Успешное создание"
// @Failure 400 {string} string "Невалидный JSON в теле запроса"
// @Failure 401 {string} string "Необходимо выполнить вход для пользователя"
// @Failure 500 {string} string "Ошибка при создании пользователя"
// @Router /tasks [post]
// @Security Bearer
func (c *Controller) CreateTask(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.CreateTask"

	// Проверяем метод запроса (должен быть POST)
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Проверяем токен и достаем user_id из него
	userID, err := validateTokenAndGetUserID(r)
	if err != nil {
		log.Printf("%s: getTokenFromRequest: %s", msgPrefix, err)
		http.Error(w, "invalid Bearer Token in Authorization Header", http.StatusUnauthorized)
		return
	}

	// Декодируем JSON-данные из тела запроса
	taskData := &model.TaskToCreate{}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(taskData); err != nil {
		log.Printf("%s: decoder.Decode: %s", msgPrefix, err)
		http.Error(w, "invalid json in body", http.StatusBadRequest)
		return
	}

	task := &model.Task{
		UserID: userID, // Добавляем user_id для таски
		Title:  taskData.Title,
		Status: taskData.Status,
	}

	dbCtx, cancel := context.WithTimeout(r.Context(), defaultDBTimeout)
	defer cancel()

	err = c.storage.CreateTask(dbCtx, task)
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

// UpdateTask обновляет существующую таску
// @Summary Обновление таски
// @Description Обновляет существующую таску. В запросе нужно передать Bearer Token
// @Tags tasks
// @Accept json
// @Produce json
// @param task body model.TaskToUpdate true "Данные существующей таски"
// @Success 200 {string} string "Успешное создание"
// @Failure 400 {string} string "Невалидный JSON в теле запроса"
// @Failure 401 {string} string "Необходимо выполнить вход для пользователя"
// @Failure 500 {string} string "Ошибка при создании пользователя"
// @Router /tasks [put]
// @Security Bearer
func (c *Controller) UpdateTask(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.CreateTask"

	// Проверяем метод запроса (должен быть PUT)
	if r.Method != http.MethodPut {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Проверяем токен и достаем user_id из него
	userID, err := validateTokenAndGetUserID(r)
	if err != nil {
		log.Printf("%s: getTokenFromRequest: %s", msgPrefix, err)
		http.Error(w, "invalid Bearer Token in Authorization Header", http.StatusUnauthorized)
		return
	}

	// Декодируем JSON-данные из тела запроса
	taskData := &model.TaskToUpdate{}
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(taskData); err != nil {
		log.Printf("%s: decoder.Decode: %s", msgPrefix, err)
		http.Error(w, "invalid json in body", http.StatusBadRequest)
		return
	}

	task := &model.Task{
		ID:     taskData.ID,
		UserID: userID, // Добавляем user_id для таски
		Title:  taskData.Title,
		Status: taskData.Status,
	}

	dbCtx, cancel := context.WithTimeout(r.Context(), defaultDBTimeout)
	defer cancel()

	err = c.storage.UpdateTaskByID(dbCtx, task)
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

// GetTasks возвращает список тасок для пользователя
// @Summary Вывод списка тасок
// @Description Выводит список тасок. В запросе нужно передать Bearer Token
// @Tags tasks
// @Produce json
// @Success 200 {array} model.Task "Список тасок"
// @Failure 401 {string} string "Необходимо выполнить вход для пользователя"
// @Failure 500 {string} string "Ошибка при создании пользователя"
// @Router /tasks [get]
// @Security Bearer
func (c *Controller) GetTasks(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.GetTasks"

	// Проверяем метод запроса (должен быть GET)
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Проверяем токен и достаем user_id из него
	userID, err := validateTokenAndGetUserID(r)
	if err != nil {
		log.Printf("%s: getTokenFromRequest: %s", msgPrefix, err)
		http.Error(w, "invalid Bearer Token in Authorization Header", http.StatusUnauthorized)
		return
	}

	dbCtx, cancel := context.WithTimeout(r.Context(), defaultDBTimeout)
	defer cancel()

	tasks, err := c.storage.GetTasksByUserID(dbCtx, userID)
	if err != nil {
		log.Printf("%s: storage.GetTasks: %s", msgPrefix, err)
		http.Error(w, "error getting tasks", http.StatusInternalServerError)
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

// validateTokenAndGetUserID проверяет токен и возвращает user_id из него
func validateTokenAndGetUserID(r *http.Request) (uuid.UUID, error) {
	userToken, err := getTokenFromRequest(r)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("getTokenFromRequest: %w", err)
	}

	userIDStr, err := jwttoken.ValidateToken(userToken)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("jwttoken.ValidateToken: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("uuid.Parse: %w", err)
	}

	return userID, nil
}

// getTokenFromRequest возвращает Bearer токен из заголовка Authorization
func getTokenFromRequest(r *http.Request) (string, error) {
	// Получаем значение заголовка Authorization
	authHeader := r.Header.Get("Authorization")

	// Проверяем наличие заголовка Authorization
	if len(authHeader) == 0 {
		return "", fmt.Errorf("missing Authorization header")
	}

	const tokenPrefix = "Bearer "
	// Проверяем, что заголовок начинается с tokenPrefix
	if !strings.HasPrefix(authHeader, tokenPrefix) {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	// Извлекаем токен, удаляя tokenPrefix из строки
	token := strings.TrimPrefix(authHeader, tokenPrefix)

	if len(token) == 0 {
		return "", fmt.Errorf("empty token")
	}

	return token, nil
}
