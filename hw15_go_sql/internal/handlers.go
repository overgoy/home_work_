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

	json.NewEncoder(w).Encode(users)
}
