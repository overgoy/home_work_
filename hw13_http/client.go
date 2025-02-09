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

	response, reqErr := sendRequest(client, targetURL, method, requestData)
	if reqErr != nil {
		log.Printf("ошибка при выполнении запроса: %v", reqErr)
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

	parsedURL, urlErr := url.Parse(targetURL)
	if urlErr != nil {
		return nil, fmt.Errorf("ошибка парсинга url: %w", urlErr)
	}

	requestObject := &http.Request{
		Method: method,
		URL:    parsedURL,
		Body:   io.NopCloser(requestBody),
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	response, httpErr := client.Do(requestObject)
	if httpErr != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", httpErr)
	}

	return response, nil
}
