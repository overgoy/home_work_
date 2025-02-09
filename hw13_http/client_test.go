package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
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
	_, err := sendRequest(client, server.URL, "GET", "")

	if err != nil {
		t.Fatalf("Ошибка при выполнении GET-запроса: %v", err)
	}
}

func TestRunClientPOST(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockServer))
	defer server.Close()

	postData := `{"key": "value"}`

	client := &http.Client{}
	_, err := sendRequest(client, server.URL, "POST", postData)

	if err != nil {
		t.Fatalf("Ошибка при выполнении POST-запроса: %v", err)
	}
}

func TestSendRequest_InvalidURL(t *testing.T) {
	client := &http.Client{}
	_, err := sendRequest(client, "htp://invalid-url", "GET", "")

	if err == nil {
		t.Fatal("Ожидалась ошибка парсинга URL, но она не произошла")
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

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"program", server.URL, "GET"}

	RunClient(server.URL, "GET")
}

func TestRunClientInvalidMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockServer))
	defer server.Close()

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"program", server.URL, "PUT"}

	RunClient(server.URL, "PUT")
}
