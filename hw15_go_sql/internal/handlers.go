package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fixme_my_friend/hw15_go_sql/db/db"
)

type Server struct {
	Queries *db.Queries
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req db.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.CreateUser(r.Context(), req)
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

func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	user, err := s.Queries.GetUser(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Ошибка получения пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req db.UpdateUserParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.UpdateUser(r.Context(), req)
	if err != nil {
		http.Error(w, "Ошибка обновления пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = s.Queries.DeleteUser(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Ошибка удаления пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var req db.CreateProductParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.CreateProduct(r.Context(), req)
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

	json.NewEncoder(w).Encode(products)
}

func (s *Server) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	product, err := s.Queries.GetProduct(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Ошибка получения продукта", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (s *Server) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var req db.UpdateProductParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.UpdateProduct(r.Context(), req)
	if err != nil {
		http.Error(w, "Ошибка обновления продукта", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = s.Queries.DeleteProduct(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Ошибка удаления продукта", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req db.CreateOrderParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	err := s.Queries.CreateOrder(r.Context(), req)
	if err != nil {
		http.Error(w, "Ошибка создания заказа", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := s.Queries.GetOrders(r.Context())
	if err != nil {
		http.Error(w, "Ошибка получения заказов", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (s *Server) GetOrdersByUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	orders, err := s.Queries.GetOrdersByUser(r.Context(), int32(userID))
	if err != nil {
		http.Error(w, "Ошибка получения заказов пользователя", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (s *Server) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = s.Queries.DeleteOrder(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Ошибка удаления заказа", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetUserStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	stats, err := s.Queries.GetUserStatistics(r.Context(), int32(userID))
	if err != nil {
		http.Error(w, "Ошибка получения статистики", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stats)
}
