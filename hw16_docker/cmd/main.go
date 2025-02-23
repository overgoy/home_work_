package main

import (
	"fmt"
	"log"

	"github.com/fixme_my_friend/hw16_docker/db"
	"github.com/fixme_my_friend/hw16_docker/internal"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(fmt.Errorf("Ошибка подключения к БД: %w", err))
	}
	defer dbConn.Close()

	err = db.CreateTablesIfNotExist(dbConn)
	if err != nil {
		log.Fatal(fmt.Errorf("Ошибка при создании таблиц: %w", err))
	}

	internal.StartServer(dbConn)
}
