package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTables(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать mock для базы данных: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS Users").
		WillReturnResult(sqlmock.NewResult(0, 0)) // Указываем 0 для LastInsertId и RowsAffected

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS Products").
		WillReturnResult(sqlmock.NewResult(0, 0)) // Указываем 0 для LastInsertId и RowsAffected

	createTables(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Ожидания не выполнены: %v", err)
	}
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO Users").
		WithArgs("Иван", "ivan@example.com", "pass").
		WillReturnResult(sqlmock.NewResult(1, 1))

	insertUser(db, "Иван", "ivan@example.com", "pass")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestInsertProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO Products").
		WithArgs("Ноутбук", 75000.50).
		WillReturnResult(sqlmock.NewResult(1, 1))

	insertProduct(db, "Ноутбук", 75000.50)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Иван").
		AddRow("Мария")

	mock.ExpectQuery("SELECT name FROM Users").
		WillReturnRows(rows)

	users := getUsers(db)
	expected := []string{"Иван", "Мария"}

	assert.Equal(t, expected, users)
	assert.NoError(t, mock.ExpectationsWereMet())
}
