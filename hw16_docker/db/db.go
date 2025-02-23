package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "db"
	port     = 5432
	user     = "shop_user"
	password = "12345"
	dbname   = "shop_db"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, connectionErr := sql.Open("postgres", psqlInfo)
	if connectionErr != nil {
		return nil, connectionErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", pingErr)
	}

	log.Println("Успешное подключение к БД!")
	return db, nil
}
