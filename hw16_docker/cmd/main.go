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
		log.Fatalf("не удалось подключиться к БД: %v", err) // Необходимо использовать %v для отображения ошибки
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения с БД: %v", err)
		}
	}()

	err = db.CreateTablesIfNotExist(dbConn)
	if err != nil {
		log.Fatalf("не удалось создать таблицы: %v", err)
	}

	internal.StartServer(dbConn)
}
