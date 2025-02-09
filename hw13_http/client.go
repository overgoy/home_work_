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

	req, err := createRequest(ctx, url, method, data)
	if err != nil {
		log.Printf("Ошибка при создании запроса: %v", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Printf("Ошибка при чтении ответа: %v", readErr)
		return
	}

	fmt.Println(string(body))
}

func createRequest(ctx context.Context, url, method, data string) (*http.Request, error) {
	var req *http.Request
	var err error

	if method == "POST" {
		if data == "" {
			return nil, fmt.Errorf("POST метод требует передачи данных")
		}
		req, err = http.NewRequest("POST", url, strings.NewReader(data))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req.WithContext(ctx), nil
}
