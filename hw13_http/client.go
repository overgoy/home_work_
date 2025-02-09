package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func RunClient(targetURL string, method string) {
	var requestData string
	if len(os.Args) >= 4 {
		requestData = os.Args[3]
	}

	client := &http.Client{Timeout: 10 * time.Second}

	response, requestErr := sendRequest(client, targetURL, method, requestData)
	if requestErr != nil {
		log.Printf("ошибка при выполнении запроса: %v", requestErr)
		return
	}
	defer response.Body.Close()

	responseBody, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Printf("ошибка при чтении ответа: %v", readErr)
		return
	}

	fmt.Println(string(responseBody))
}

func sendRequest(client *http.Client, targetURL, method, requestData string) (*http.Response, error) {
	var requestBody *bytes.Buffer
	if method == "POST" {
		if requestData == "" {
			return nil, fmt.Errorf("post метод требует передачи данных")
		}
		requestBody = bytes.NewBuffer([]byte(requestData))
	} else {
		requestBody = bytes.NewBuffer(nil)
	}

	parsedURL, parseErr := url.Parse(targetURL)
	if parseErr != nil {
		return nil, fmt.Errorf("ошибка парсинга url: %w", parseErr)
	}

	requestObject := &http.Request{
		Method: method,
		URL:    parsedURL,
		Body:   io.NopCloser(requestBody),
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	response, responseErr := client.Do(requestObject)
	if responseErr != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", responseErr)
	}

	return response, nil
}
