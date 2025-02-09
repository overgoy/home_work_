package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func RunClient(url string, method string) {
	var data string
	if len(os.Args) >= 4 {
		data = os.Args[3]
	}

	req, reqErr := createRequest(url, method, data)
	if reqErr != nil {
		log.Printf("Ошибка при создании запроса: %v", reqErr)
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}

	resp, respErr := client.Do(req)
	if respErr != nil {
		log.Printf("Ошибка при выполнении запроса: %v", respErr)
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

func createRequest(url, method, data string) (*http.Request, error) {
	if method == "POST" {
		if data == "" {
			return nil, fmt.Errorf("POST метод требует передачи данных")
		}
		req, err := http.NewRequest("POST", url, strings.NewReader(data))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		return req, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
