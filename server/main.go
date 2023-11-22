package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
todo:
- подключить бд, сохранять и подгружать данные
- сделать регистрацию и вход
- таски для каждого пользователя свои
- когда создаем новую таску или меняем статус текущей, то вызывать метод в API, чтобы записать обновленные данные
*/

func main() {
	// Обработчик запросов
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		// Создаем список тасок
		tasks := []Task{
			{
				Status: "InProgress",
				Title:  "Привет",
			},
			{
				Status: "ToDo",
				Title:  "Пока",
			},
			{
				Status: "Done",
				Title:  "1",
			},
			{
				Status: "Deleted",
				Title:  "2",
			},
		}

		// Преобразуем данные в JSON
		jsonData, err := json.Marshal(tasks)
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
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
