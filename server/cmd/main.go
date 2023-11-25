package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/l-orlov/simple-todo-list/server/internal/controller"
	"github.com/l-orlov/simple-todo-list/server/internal/server"
	"github.com/l-orlov/simple-todo-list/server/internal/store"
	"github.com/rs/cors"
)

func init() {
	// Используем таймзону UTC по дефолту для временных меток
	time.Local = time.UTC
}

// @title Your API Title
// @version 1.0
// @description Your API description. You can use Markdown here.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	ctx := context.Background()

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
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

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
	port := "8080"
	srv := server.New(port, c.Handler(r))
	go func() {
		if err = srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error while running http server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	log.Print("Server shutting down")

	if err = srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down: %v", err)
	}
}
