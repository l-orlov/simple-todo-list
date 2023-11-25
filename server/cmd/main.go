package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/l-orlov/simple-todo-list/server/internal/controller"
	"github.com/l-orlov/simple-todo-list/server/internal/store"
	"github.com/rs/cors"
)

func init() {
	// Используем таймзону UTC по дефолту для временных меток
	time.Local = time.UTC
}

/*
todo:
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
	r.HandleFunc("/register", routsController.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/login", routsController.LoginUser).Methods(http.MethodPost)
	r.HandleFunc("/tasks", routsController.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", routsController.UpdateTask).Methods(http.MethodPut)
	r.HandleFunc("/tasks", routsController.GetTasks).Methods(http.MethodGet)

	// Создаем экземпляр CORS с настройками по умолчанию
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Запуск веб-сервера на порту 8080 с поддержкой CORS
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	handler := cors.AllowAll().Handler(r)
	http.Handle("/", c.Handler(handler))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
