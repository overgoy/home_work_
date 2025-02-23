package main

import (
	"log"

	"github.com/fixme_my_friend/hw16_docker/db"
	"github.com/fixme_my_friend/hw16_docker/internal"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("не удалось подключиться к БД: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("ошибка при закрытии соединения с БД: %v", err)
		}
	}()

	if err := db.CreateTablesIfNotExist(dbConn); err != nil {
		log.Fatalf("не удалось создать таблицы: %w", err)
	}

	internal.StartServer(dbConn)
}
