package main

import (
	"log"

	"github.com/fixme_my_friend/hw15_go_sql/db"
	"github.com/fixme_my_friend/hw15_go_sql/internal"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer dbConn.Close()

	internal.StartServer(dbConn)
}
