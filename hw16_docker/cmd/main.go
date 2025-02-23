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
		closeErr := dbConn.Close()
		if closeErr != nil {
			log.Printf("ошибка при закрытии соединения с БД: %v", closeErr)
		}
	}()

	err = db.CreateTablesIfNotExist(dbConn)
	if err != nil {
		log.Fatalf("не удалось создать таблицы: %v", err)
	}

	internal.StartServer(dbConn)
}
