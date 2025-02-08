package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
