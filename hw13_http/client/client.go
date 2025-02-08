package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Использование: %s <server_url> <method> [data]", os.Args[0])
	}

	url := os.Args[1]
	method := os.Args[2]

	var req *http.Request
	var err error

	if method == "POST" {
		if len(os.Args) < 4 {
			log.Fatal("POST метод требует передачи данных")
		}
		data := os.Args[3]
		req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		log.Fatalf("Ошибка при создании запроса: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	fmt.Println(string(body))
}
