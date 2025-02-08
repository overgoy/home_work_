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

// Тестируем обработку GET-запроса
func TestHandlerGET(t *testing.T) {
	// Создаем GET-запрос
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Ошибка создания GET-запроса:", err)
	}

	// Создаем recorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler) // передаем обработчик
	handler.ServeHTTP(rr, req)           // обрабатываем запрос

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код состояния: получили %v, ожидался %v", status, http.StatusOK)
	}

	// Декодируем ответ в структуру
	var response ResponseData
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal("Ошибка декодирования JSON-ответа:", err)
	}

	// Проверяем правильность сообщения
	expectedMessage := "Получен GET-запрос для /"
	if response.Message != expectedMessage {
		t.Errorf("Неверное сообщение: получили %v, ожидалось %v", response.Message, expectedMessage)
	}
}

// Тестируем обработку POST-запроса
func TestHandlerPOST(t *testing.T) {
	// Создаем тело POST-запроса
	body := bytes.NewBufferString("тестовые данные")

	// Создаем POST-запрос
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal("Ошибка создания POST-запроса:", err)
	}

	// Устанавливаем Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Создаем recorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler) // передаем обработчик
	handler.ServeHTTP(rr, req)           // обрабатываем запрос

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный код состояния: получили %v, ожидался %v", status, http.StatusOK)
	}

	// Декодируем ответ в структуру
	var response ResponseData
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal("Ошибка декодирования JSON-ответа:", err)
	}

	// Проверяем правильность сообщения
	expectedMessage := "Получен POST-запрос с телом: тестовые данные"
	if response.Message != expectedMessage {
		t.Errorf("Неверное сообщение: получили %v, ожидалось %v", response.Message, expectedMessage)
	}
}
