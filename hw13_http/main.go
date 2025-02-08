package main

import (
	"log"
	"os"
)

func main() {
	// Проверка аргументов и запуск клиента или сервера
	if len(os.Args) < 3 {
		log.Println("Использование: <server_url> <method> [data]")
		return
	}

	url := os.Args[1]
	method := os.Args[2]

	// Запуск клиента (для POST/GET запросов)
	RunClient(url, method)
}
