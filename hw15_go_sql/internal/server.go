package internal

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/fixme_my_friend/hw15_go_sql/db/db"
	"github.com/go-chi/chi/v5"
)

func StartServer(dbConn *sql.DB) {
	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	r := chi.NewRouter()
	r.Post("/users", server.CreateUserHandler)
	r.Get("/users", server.GetUsersHandler)

	addr := ":8080"
	log.Printf("Сервер запущен на %s", addr)

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
