package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/l-orlov/simple-todo-list/server/internal/jwttoken"
	"github.com/l-orlov/simple-todo-list/server/internal/model"
	"github.com/l-orlov/simple-todo-list/server/internal/store"
)

// defaultDBTimeout - дефолтный таймаут для запросов в БД
const defaultDBTimeout = 5 * time.Second

// RegisterUser регистрирует нового пользователя.
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя с переданными данными в запросе
// @Tags register
// @Accept json
// @Param task body model.UserLoginData true "Email и пароль для пользователя"
// @Success 200 {string} string "Пользователь успешно зарегистрирован"
// @Failure 400 {string} string "Невалидный JSON в теле запроса"
// @Failure 409 {string} string "Пользователь уже существует"
// @Failure 500 {string} string "Ошибка при создании пользователя"
// @Router /register [post]
func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.RegisterUser"

	// Проверяем метод запроса (должен быть POST)
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON-данные из тела запроса
	userData := &model.UserLoginData{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(userData); err != nil {
		log.Printf("%s: decoder.Decode: %s", msgPrefix, err)
		http.Error(w, "invalid json in body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := model.HashPassword(userData.Password)
	if err != nil {
		log.Printf("%s: model.HashPassword: %s", msgPrefix, err)
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Email:    userData.Email,
		Password: hashedPassword,
	}

	dbCtx, cancel := context.WithTimeout(r.Context(), defaultDBTimeout)
	defer cancel()

	err = c.storage.CreateUser(dbCtx, user)
	if err != nil {
		log.Printf("%s: storage.CreateUser: %s", msgPrefix, err)
		if errors.Is(err, store.ErrViolatesUniqConst) {
			http.Error(w, "user already exists", http.StatusConflict)
			return
		}

		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	// При успехе возвращаем пустую структуру
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&struct{}{}); err != nil {
		log.Printf("%s: json encode response: %s", msgPrefix, err)
		http.Error(w, "error encoding json response", http.StatusInternalServerError)
		return
	}
}

// LoginUser выполняет вход в систему для пользователя
// @Summary Вход пользователя
// @Description Выполняет вход в систему для пользователя с переданными данными в запросе
// @Tags login
// @Accept  json
// @Produce  json
// @Param task body model.UserLoginData true "Email и пароль для пользователя"
// @Success 200 {string} string "Успешный вход в систему"
// @Failure 400 {string} string "Невалидный JSON в теле запроса"
// @Failure 401 {string} string "Неверный пароль"
// @Failure 404 {string} string "Пользователь не найден"
// @Failure 500 {string} string "Ошибка при создании пользователя"
// @Router /login [post]
func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	msgPrefix := "handler.CreateTask"

	// Проверяем метод запроса (должен быть POST)
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON-данные из тела запроса
	userData := &model.UserLoginData{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(userData); err != nil {
		log.Printf("%s: decoder.Decode: %s", msgPrefix, err)
		http.Error(w, "invalid json in body", http.StatusBadRequest)
		return
	}

	dbCtx, cancel := context.WithTimeout(r.Context(), defaultDBTimeout)
	defer cancel()

	user, err := c.storage.GetUserByEmail(dbCtx, userData.Email)
	if err != nil {
		log.Printf("%s: storage.GetUserByEmail: %s", msgPrefix, err)
		if errors.Is(err, store.ErrNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		http.Error(w, "error getting user", http.StatusInternalServerError)
		return
	}

	if !model.CheckPasswordHash(user.Password, userData.Password) {
		log.Printf("%s: wrong password: %s", msgPrefix, err)
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	userToken, err := jwttoken.GenerateToken(user.ID.String())
	if err != nil {
		log.Printf("%s: token.GenerateToken: %s", msgPrefix, err)
		http.Error(w, "error generating token for user", http.StatusInternalServerError)
		return
	}

	userTokenResp := &model.UserToken{
		Token: userToken,
	}

	// Кодируем JSON и отправляем в ответе
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(&userTokenResp); err != nil {
		log.Printf("%s: json encode response: %s", msgPrefix, err)
		http.Error(w, "error encoding json response", http.StatusInternalServerError)
		return
	}
}
