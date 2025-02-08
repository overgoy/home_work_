package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен %s запрос для %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	var response Response

	// Обработка POST-запроса
	if r.Method == http.MethodPost {
		body, readErr := io.ReadAll(r.Body)
		if readErr != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusInternalServerError)
			return
		}
		response = Response{Message: fmt.Sprintf("Получен POST-запрос с телом: %s", string(body))}
	} else {
		// Обработка GET-запроса
		response = Response{Message: fmt.Sprintf("Получен GET-запрос для %s", r.URL.Path)}
	}

	// Отправляем JSON-ответ
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		http.Error(w, "Ошибка при отправке ответа", http.StatusInternalServerError)
	}
}

func mainServer() {
	// Проверка аргументов командной строки
	if len(os.Args) < 3 {
		log.Fatalf("Использование: %s <адрес> <порт>", os.Args[0])
	}
	addr := os.Args[1]
	port := os.Args[2]

	// Регистрируем обработчик
	http.HandleFunc("/", handler)
	log.Printf("Запуск сервера на %s:%s", addr, port)

	// Настроим сервер с тайм-аутами
	srv := &http.Server{
		Addr:         addr + ":" + port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	// Запуск сервера
	log.Printf("Сервер запущен на %s:%s", addr, port)
	serverErr := srv.ListenAndServe()
	if serverErr != nil {
		log.Fatalf("Ошибка запуска сервера: %v", serverErr)
	}
}
