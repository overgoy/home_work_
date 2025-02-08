package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func startServer() {
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
