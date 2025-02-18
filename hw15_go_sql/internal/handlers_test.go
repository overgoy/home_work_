package internal

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/fixme_my_friend/hw15_go_sql/db"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	dbConn, err := db.Connect()
	if err != nil {
		t.Fatal("Ошибка подключения к тестовой БД:", err)
	}
	return dbConn
}

func TestCreateUserHandler(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	user := map[string]string{
		"name":     "Иван",
		"email":    "ivan@example.com",
		"password": "pass123",
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.CreateUserHandler(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetUsersHandler(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	server.GetUsersHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateProductHandler(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	product := map[string]interface{}{
		"name":  "Ноутбук",
		"price": 999.99,
	}
	body, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	server.CreateProductHandler(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetProductsHandler(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	req, _ := http.NewRequest("GET", "/products", nil)
	rr := httptest.NewRecorder()

	server.GetProductsHandler(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
