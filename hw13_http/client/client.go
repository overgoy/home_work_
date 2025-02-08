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

func main() {
	// Проверяем количество аргументов
	if len(os.Args) < 3 {
		log.Println("Использование: <server_url> <method> [data]")
		return
	}

	// Получаем URL и метод
	url := os.Args[1]
	method := os.Args[2]

	// Переменная для данных в случае метода POST
	var data string
	if len(os.Args) >= 4 {
		data = os.Args[3]
	}

	// Создаем контекст для запроса
	ctx := context.Background()

	// Переменная для запроса
	var req *http.Request
	var reqCreateErr error

	// Обработка метода POST
	if method == "POST" {
		if data == "" {
			log.Println("POST метод требует передачи данных")
			return
		}
		req, reqCreateErr = http.NewRequestWithContext(ctx, method, url, nil)
		if reqCreateErr != nil {
			log.Printf("Ошибка при создании POST-запроса: %v", reqCreateErr)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		// Устанавливаем тело запроса с данными
		req.Body = io.NopCloser(strings.NewReader(data))
	} else {
		// Обработка метода GET
		req, reqCreateErr = http.NewRequestWithContext(ctx, "GET", url, nil)
		if reqCreateErr != nil {
			log.Printf("Ошибка при создании GET-запроса: %v", reqCreateErr)
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
