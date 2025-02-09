package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type ResponseData struct {
	Message string `json:"message"`
}

func TestHandlerGET(t *testing.T) {
	req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/"},
		Body:   nil,
		Header: map[string][]string{},
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

	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Path: "/"},
		Body:   io.NopCloser(body),
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
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

	expectedMessage := "Получен POST-запрос с телом: тестовые данные"
	if response.Message != expectedMessage {
		t.Errorf("Неверное сообщение: получили %v, ожидалось %v", response.Message, expectedMessage)
	}
}
