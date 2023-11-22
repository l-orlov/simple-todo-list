package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/l-orlov/simple-todo-list/server/store"
)

/*
todo:
- сохранять данные в БД
- сделать регистрацию и вход
- таски для каждого пользователя свои
- когда создаем новую таску или меняем статус текущей, то вызывать метод в API, чтобы записать обновленные данные
*/

func main() {
	ctx := context.Background()

	// Connect to DB
	storage, err := store.New()
	if err != nil {
		log.Fatalf("store.New: %s", err)
	}

	initialTasks, err := storage.GetTasks(ctx)
	if err != nil {
		log.Fatalf("storage.GetTasks: %s", err)
	}

	// Обработчик запросов
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		// Преобразуем данные в JSON
		jsonData, err := json.Marshal(initialTasks)
		if err != nil {
			http.Error(w, "error with json marshalling", http.StatusInternalServerError)
			return
		}

		// Устанавливаем заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")
		// Устанавливаем заголовок Access-Control-Allow-Origin для разрешения запросов с любых доменов
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Устанавливаем статус 200
		w.WriteHeader(http.StatusOK)

		// Отправляем JSON в ответе
		_, _ = w.Write(jsonData)
	})

	http.HandleFunc("/create-task", func(w http.ResponseWriter, r *http.Request) {
		// Преобразуем данные в JSON
		jsonData, err := json.Marshal(initialTasks)
		if err != nil {
			http.Error(w, "error with json marshalling", http.StatusInternalServerError)
			return
		}

		// Устанавливаем заголовок Content-Type
		w.Header().Set("Content-Type", "application/json")
		// Устанавливаем заголовок Access-Control-Allow-Origin для разрешения запросов с любых доменов
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Устанавливаем статус 200
		w.WriteHeader(http.StatusOK)

		// Отправляем JSON в ответе
		_, _ = w.Write(jsonData)
	})

	// Слушаем порт 8080
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
