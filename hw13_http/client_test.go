package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"message": "GET-запрос выполнен"}`)
	case "POST":
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close() // Закрываем тело запроса
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprintf(`{"message": "POST-запрос выполнен с данными: %s"}`, string(body)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, `{"message": "Метод не поддерживается"}`)
	}
}

func TestRunClientGET(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockServer))
	defer server.Close()

	client := &http.Client{}
	resp, err := sendRequest(client, server.URL, "GET", "")
	if err != nil {
		t.Fatalf("Ошибка при выполнении GET-запроса: %v", err)
	}
	defer resp.Body.Close() // Закрываем тело ответа
}

func TestRunClientPOST(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockServer))
	defer server.Close()

	client := &http.Client{}
	resp, err := sendRequest(client, server.URL, "POST", `{"key": "value"}`)
	if err != nil {
		t.Fatalf("Ошибка при выполнении POST-запроса: %v", err)
	}
	defer resp.Body.Close() // Закрываем тело ответа
}

func TestSendRequest_InvalidURL(t *testing.T) {
	client := &http.Client{}
	_, err := sendRequest(client, "htp://invalid-url", "GET", "")

	if err == nil {
		t.Fatal("Ожидалась ошибка парсинга URL, но её не произошло")
	}
}

func TestSendRequest_NoDataForPOST(t *testing.T) {
	client := &http.Client{}
	_, err := sendRequest(client, "http://localhost", "POST", "")

	if err == nil {
		t.Fatal("Ожидалась ошибка: POST метод требует передачи данных, но ошибка не произошла")
	}
}

func TestRunClientGETSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockServer))
	defer server.Close()

	RunClient(server.URL, "GET")
}

func TestRunClientInvalidMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockServer))
	defer server.Close()

	RunClient(server.URL, "PUT")
}
