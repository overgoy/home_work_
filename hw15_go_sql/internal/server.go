package internal

import (
	"database/sql"
	"github.com/fixme_my_friend/hw15_go_sql/db/db"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func StartServer(dbConn *sql.DB) {
	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	r := chi.NewRouter()
	r.Post("/users", server.CreateUserHandler)
	r.Get("/users", server.GetUsersHandler)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
