package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Получен %s запрос для %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	var response Response

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusInternalServerError)
			return
		}
		response = Response{Message: fmt.Sprintf("Получен POST-запрос с телом: %s", string(body))}
	} else {
		response = Response{Message: fmt.Sprintf("Получен GET-запрос для %s", r.URL.Path)}
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Использование: %s <адрес> <порт>", os.Args[0])
	}
	addr := os.Args[1]
	port := os.Args[2]

	http.HandleFunc("/", handler)
	log.Printf("Запуск сервера на %s:%s", addr, port)
	log.Fatal(http.ListenAndServe(addr+":"+port, nil))
}
