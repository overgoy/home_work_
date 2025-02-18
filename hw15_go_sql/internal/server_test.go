package internal

import (
	"database/sql"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	dbConn, err := sql.Open("postgres", "host=localhost port=5432 user=shop_user password=12345 dbname=shop_db sslmode=disable")
	if err != nil {
		t.Fatal("Ошибка подключения к тестовой БД:", err)
	}
	defer dbConn.Close()

	go StartServer(dbConn)

	resp, err := http.Get("http://localhost:8080/users")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
