package main

import (
	"fmt"
	"log"

	"github.com/fixme_my_friend/hw16_docker/db"
	"github.com/fixme_my_friend/hw16_docker/internal"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, connErr := db.Connect()
	if connErr != nil {
		log.Fatal(fmt.Errorf("ошибка подключения к БД: %w", connErr))
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения с БД: %v", err)
		}
	}()

	internal.StartServer(dbConn)
}
