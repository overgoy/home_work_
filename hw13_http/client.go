package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func RunClient(url string, method string) {
	var requestData string
	if len(os.Args) >= 4 {
		requestData = os.Args[3]
	}

	client := &http.Client{Timeout: 10 * time.Second}

	response, requestError := executeRequest(client, url, method, requestData)
	if requestError != nil {
		log.Printf("Ошибка при выполнении запроса: %v", requestError)
		return
	}
	defer response.Body.Close()

	responseBody, readError := io.ReadAll(response.Body)
	if readError != nil {
		log.Printf("Ошибка при чтении ответа: %v", readError)
		return
	}

	fmt.Println(string(responseBody))
}

func executeRequest(client *http.Client, url, method, requestData string) (*http.Response, error) {
	var requestBody *bytes.Reader
	if method == "POST" {
		if requestData == "" {
			return nil, fmt.Errorf("POST метод требует передачи данных")
		}
		requestBody = bytes.NewReader([]byte(requestData))
	} else {
		requestBody = bytes.NewReader(nil)
	}

	requestObject, requestCreationError := client.Post(url, "application/json", requestBody)
	if requestCreationError != nil {
		return nil, requestCreationError
	}

	return requestObject, nil
}
