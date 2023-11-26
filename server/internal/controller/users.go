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

// RegisterUser регистрирует нового пользователя
// @summary Регистрирует нового пользователя
// @description Регистрирует нового пользователя с использованием данных из тела запроса.
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "Username"
// @Param   password     body    string     true        "Password"
// @Success 200 {string} string  "ok"
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

	// При успехе ничего не возвращаем. В ответе будет статус 200 Ok
}

// LoginUser входит пользователя в систему
// @summary Вход пользователя
// @description Входит пользователя в систему с использованием данных из тела запроса.
// @Accept  json
// @Produce  json
// @Param   username     body    string     true        "Username"
// @Param   password     body    string     true        "Password"
// @Success 200 {string} string  "ok"
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
