package main

import (
	"github.com/fixme_my_friend/hw15_go_sql/db"
	"github.com/fixme_my_friend/hw15_go_sql/internal"
	"log"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer dbConn.Close()

	internal.StartServer(dbConn)
}
