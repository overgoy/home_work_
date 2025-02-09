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

	response, requestExecutionError := sendRequest(client, targetURL, method, requestData)
	if requestExecutionError != nil {
		log.Printf("Ошибка при выполнении запроса: %v", requestExecutionError)
		return
	}
	defer response.Body.Close()

	responseBody, responseReadError := io.ReadAll(response.Body)
	if responseReadError != nil {
		log.Printf("Ошибка при чтении ответа: %v", responseReadError)
		return
	}

	fmt.Println(string(responseBody))
}

func sendRequest(client *http.Client, targetURL, method, requestData string) (*http.Response, error) {
	var requestBody *bytes.Buffer
	if method == "POST" {
		if requestData == "" {
			return nil, fmt.Errorf("POST метод требует передачи данных")
		}
		requestBody = bytes.NewBuffer([]byte(requestData))
	} else {
		requestBody = bytes.NewBuffer(nil)
	}

	parsedURL, parseErr := url.Parse(targetURL)
	if parseErr != nil {
		return nil, fmt.Errorf("Ошибка парсинга URL: %w", parseErr)
	}

	requestObject := &http.Request{
		Method: method,
		URL:    parsedURL,
		Body:   io.NopCloser(requestBody),
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	response, executionError := client.Do(requestObject)
	if executionError != nil {
		return nil, executionError
	}

	return response, nil
}
