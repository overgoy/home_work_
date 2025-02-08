package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ResponseData struct {
	Message string `json:"message"`
}

func TestHandlerGET(t *testing.T) {
	req, reqErr := http.NewRequest("GET", "/", nil)
	if reqErr != nil {
		t.Fatal("Ошибка создания GET-запроса")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код состояния: получили %v, ожидался %v", status, http.StatusOK)
	}

	var response ResponseData
	if decodeErr := json.NewDecoder(rr.Body).Decode(&response); decodeErr != nil {
		t.Fatal("Ошибка декодирования JSON-ответа")
	}
}

func TestHandlerPOST(t *testing.T) {
	body := bytes.NewBufferString("тестовые данные")
	req, reqErr := http.NewRequest("POST", "/", body)
	if reqErr != nil {
		t.Fatal("Ошибка создания POST-запроса")
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код состояния: получили %v, ожидался %v", status, http.StatusOK)
	}

	var response ResponseData
	if decodeErr := json.NewDecoder(rr.Body).Decode(&response); decodeErr != nil {
		t.Fatal("Ошибка декодирования JSON-ответа")
	}
}
