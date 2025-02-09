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
		t.Fatal("Ошибка создания GET-запроса:", reqErr)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код состояния: получили %v, ожидался %v", status, http.StatusOK)
	}

	var response ResponseData
	decodeErr := json.NewDecoder(rr.Body).Decode(&response)
	if decodeErr != nil {
		t.Fatal("Ошибка декодирования JSON-ответа:", decodeErr)
	}

	expectedMessage := "Получен GET-запрос для /"
	if response.Message != expectedMessage {
		t.Errorf("Неверное сообщение: получили %v, ожидалось %v", response.Message, expectedMessage)
	}
}

func TestHandlerPOST(t *testing.T) {
	body := bytes.NewBufferString("тестовые данные")

	req, reqErr := http.NewRequest("POST", "/", body)
	if reqErr != nil {
		t.Fatal("Ошибка создания POST-запроса:", reqErr)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код состояния: получили %v, ожидался %v", status, http.StatusOK)
	}

	var response ResponseData
	decodeErr := json.NewDecoder(rr.Body).Decode(&response)
	if decodeErr != nil {
		t.Fatal("Ошибка декодирования JSON-ответа:", decodeErr)
	}

	expectedMessage := "Получен POST-запрос с телом: тестовые данные"
	if response.Message != expectedMessage {
		t.Errorf("Неверное сообщение: получили %v, ожидалось %v", response.Message, expectedMessage)
	}
}
