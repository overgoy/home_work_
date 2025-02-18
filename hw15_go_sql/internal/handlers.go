package internal

import (
	"encoding/json"
	"net/http"

	"github.com/fixme_my_friend/hw15_go_sql/db/db"
)

type Server struct {
	Queries *db.Queries
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.CreateUser(r.Context(), db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := s.Queries.GetUsers(r.Context())
	if err != nil {
		http.Error(w, "Ошибка получения пользователей", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Ошибка кодирования списка пользователей", http.StatusInternalServerError)
		return
	}
}

func (s *Server) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.CreateProduct(r.Context(), db.CreateProductParams{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		http.Error(w, "Ошибка создания продукта", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := s.Queries.GetProducts(r.Context())
	if err != nil {
		http.Error(w, "Ошибка получения продуктов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Ошибка кодирования списка продуктов", http.StatusInternalServerError)
		return
	}
}
