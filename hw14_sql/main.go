package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shop_user"
	password = "12345"
	dbname   = "shop_db"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}
	fmt.Println("✅ Успешное подключение к БД!")

	createTables(db)

	insertUser(db, "Иван", "ivan@example.com", "pass")
	insertProduct(db, "Ноутбук", 75000.50)

	users := getUsers(db)
	fmt.Println("📌 Список пользователей:", users)
}

func createTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS Users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS Products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблиц:", err)
	}
	fmt.Println("✅ Таблицы успешно созданы!")
}

func insertUser(db *sql.DB, name, email, password string) {
	_, err := db.Exec("INSERT INTO Users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
	if err != nil {
		log.Fatal("Ошибка вставки пользователя:", err)
	}
	fmt.Println("✅ Пользователь добавлен:", name)
}

func insertProduct(db *sql.DB, name string, price float64) {
	_, err := db.Exec("INSERT INTO Products (name, price) VALUES ($1, $2)", name, price)
	if err != nil {
		log.Fatal("Ошибка вставки товара:", err)
	}
	fmt.Println("✅ Товар добавлен:", name)
}

func getUsers(db *sql.DB) []string {
	rows, err := db.Query("SELECT name FROM Users")
	if err != nil {
		log.Fatal("Ошибка выборки пользователей:", err)
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		users = append(users, name)
	}
	return users
}
