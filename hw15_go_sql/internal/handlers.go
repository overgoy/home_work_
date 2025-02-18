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
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	createUserErr := s.Queries.CreateUser(r.Context(), db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if createUserErr != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, getUsersErr := s.Queries.GetUsers(r.Context())
	if getUsersErr != nil {
		http.Error(w, "Ошибка получения пользователей", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(users); encodeErr != nil {
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
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	createProductErr := s.Queries.CreateProduct(r.Context(), db.CreateProductParams{
		Name:  req.Name,
		Price: req.Price,
	})
	if createProductErr != nil {
		http.Error(w, "Ошибка создания продукта", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, getProductsErr := s.Queries.GetProducts(r.Context())
	if getProductsErr != nil {
		http.Error(w, "Ошибка получения продуктов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if encodeErr := json.NewEncoder(w).Encode(products); encodeErr != nil {
		http.Error(w, "Ошибка кодирования списка продуктов", http.StatusInternalServerError)
		return
	}
}
