package main

import (
	"log"
	"os"
)

func main() {
	// Проверка аргументов командной строки
	if len(os.Args) < 3 {
		log.Fatalf("Использование: %s <адрес> <порт>", os.Args[0])
	}
	startServer()
}
