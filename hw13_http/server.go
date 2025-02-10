package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен %s запрос для %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	var response Response

	switch r.Method {
	case http.MethodGet:
		response = Response{Message: fmt.Sprintf("Получен GET-запрос для %s", r.URL.Path)}
	case http.MethodPost:
		body, rErr := io.ReadAll(r.Body)
		if rErr != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusInternalServerError)
			return
		}
		response = Response{Message: fmt.Sprintf("Получен POST-запрос с телом: %s", string(body))}
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		http.Error(w, "Ошибка при отправке ответа", http.StatusInternalServerError)
	}
}

func startServer(addr string, port string) {
	http.HandleFunc("/", handler)
	log.Printf("Запуск сервера на %s:%s", addr, port)

	srv := &http.Server{
		Addr:         addr + ":" + port,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Printf("Сервер запущен на %s:%s", addr, port)
	serverErr := srv.ListenAndServe()
	if serverErr != nil {
		log.Printf("Ошибка запуска сервера: %v", serverErr)
	}
}
