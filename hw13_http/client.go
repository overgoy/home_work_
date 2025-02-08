package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// RunClient — функция для выполнения HTTP-запросов.
func RunClient(url string, method string) {
	// Переменная для данных в случае метода POST
	var data string
	if len(os.Args) >= 4 {
		data = os.Args[3]
	}

	// Создаем контекст для запроса
	ctx := context.Background()

	// Переменная для запроса
	var req *http.Request
	var reqErr error

	// Обработка метода POST
	if method == "POST" {
		if data == "" {
			log.Println("POST метод требует передачи данных")
			return
		}
		req, reqErr = http.NewRequestWithContext(ctx, method, url, nil)
		if reqErr != nil {
			log.Printf("Ошибка при создании POST-запроса: %v", reqErr)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		// Устанавливаем тело запроса с данными
		req.Body = io.NopCloser(strings.NewReader(data))
	} else {
		// Обработка метода GET
		req, reqErr = http.NewRequestWithContext(ctx, "GET", url, nil)
		if reqErr != nil {
			log.Printf("Ошибка при создании GET-запроса: %v", reqErr)
			return
		}
	}

	// Отправка запроса
	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Printf("Ошибка при выполнении запроса: %v", respErr)
		return
	}
	defer resp.Body.Close() // defer работает даже с return, так как он будет выполнен до выхода из main()

	// Чтение ответа
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Printf("Ошибка при чтении ответа: %v", readErr)
		return
	}

	// Выводим тело ответа
	fmt.Println(string(body))
}
