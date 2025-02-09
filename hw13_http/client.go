package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func RunClient(url string, method string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data string
	if len(os.Args) >= 4 {
		data = os.Args[3]
	}

	var req *http.Request
	var err error

	if method == "POST" {
		if data == "" {
			log.Println("Ошибка: POST метод требует передачи данных")
			return
		}
		req, err = http.NewRequest("POST", url, strings.NewReader(data))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		log.Printf("Ошибка при создании запроса: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	req = req.WithContext(ctx)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка при чтении ответа: %v", err)
		return
	}

	fmt.Println(string(body))
}
