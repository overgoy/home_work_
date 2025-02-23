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
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer dbConn.Close()

	err = db.CreateTablesIfNotExist(dbConn)
	if err != nil {
		log.Fatal("Ошибка при создании таблиц:", err)
	}

	internal.StartServer(dbConn)
}
