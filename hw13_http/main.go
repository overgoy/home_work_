package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Использование: <server_url> <method> [data]")
		return
	}

	if os.Args[1] == "server" {
		if len(os.Args) < 4 {
			log.Println("Использование: server <адрес> <порт>")
			return
		}
		addr := os.Args[2]
		port := os.Args[3]
		startServer(addr, port)
	} else {
		url := os.Args[1]
		method := os.Args[2]
		RunClient(url, method)
	}
}
