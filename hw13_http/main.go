package main

import (
	"log"
	"os"
)

// main — точка входа программы, которая запускает сервер или клиента.
func main() {
	if len(os.Args) < 3 {
		log.Println("Использование: <server_url> <method> [data]")
		return
	}

	// Если мы передаем адрес и порт, запускаем сервер
	if os.Args[1] == "server" {
		if len(os.Args) < 4 {
			log.Println("Использование: server <адрес> <порт>")
			return
		}
		addr := os.Args[2]
		port := os.Args[3]
		startServer(addr, port)
	} else {
		// Запускаем клиента
		url := os.Args[1]
		method := os.Args[2]
		RunClient(url, method)
	}
}
