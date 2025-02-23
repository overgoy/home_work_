package internal

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/fixme_my_friend/hw16_docker/db/db"
	"github.com/go-chi/chi/v5"
)

func StartServer(dbConn *sql.DB) {
	queries := db.New(dbConn)
	server := &Server{Queries: queries}

	r := chi.NewRouter()

	r.Post("/users", server.CreateUserHandler)
	r.Get("/users", server.GetUsersHandler)
	r.Get("/users/{id}", server.GetUserHandler)
	r.Put("/users/{id}", server.UpdateUserHandler)
	r.Delete("/users/{id}", server.DeleteUserHandler)
	r.Post("/products", server.CreateProductHandler)
	r.Get("/products", server.GetProductsHandler)
	r.Get("/products/{id}", server.GetProductHandler)
	r.Put("/products/{id}", server.UpdateProductHandler)
	r.Delete("/products/{id}", server.DeleteProductHandler)
	r.Post("/orders", server.CreateOrderHandler)
	r.Get("/orders", server.GetOrdersHandler)
	r.Get("/orders/user/{user_id}", server.GetOrdersByUserHandler)
	r.Delete("/orders/{id}", server.DeleteOrderHandler)
	r.Get("/users/{id}/stats", server.GetUserStatisticsHandler)

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
