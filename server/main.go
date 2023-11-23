package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l-orlov/simple-todo-list/server/controller"
	"github.com/l-orlov/simple-todo-list/server/store"
	"github.com/rs/cors"
)

/*
todo:
- сохранять данные в БД
  - обновление таски по ID
- сделать регистрацию и вход
- таски для каждого пользователя свои
- когда создаем новую таску или меняем статус текущей, то вызывать метод в API, чтобы записать обновленные данные
- добавить swagger
- добавить graceful shutdown
*/

// @title Your API Title
// @version 1.0
// @description Your API description. You can use Markdown here.
// @host localhost:8080
// @BasePath /v1
func main() {
	ctx := context.Background()
	_ = ctx

	// Connect to DB
	storage, err := store.New()
	if err != nil {
		log.Fatalf("store.New: %s", err)
	}

	// Init handler
	routsController, err := controller.New(storage)
	if err != nil {
		log.Fatalf("controller.New: %s", err)
	}

	// Инициализация маршрутизатора Gorilla Mux
	r := mux.NewRouter()

	// Регистрируем хэндлеры
	r.HandleFunc("/tasks/create", routsController.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks/get", routsController.GetTasks).Methods(http.MethodGet)

	// Создаем экземпляр CORS с настройками по умолчанию
	c := cors.Default()

	// Запуск веб-сервера на порту 8080 с поддержкой CORS
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	handler := cors.Default().Handler(r)
	http.Handle("/", c.Handler(handler))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
