package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/overgoy/home_work_/hw09_serialize/bookpb"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   float64 `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   float64 `json:"size"`
		Rate   float64 `json:"rate"`
	}{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   float64 `json:"size"`
		Rate   float64 `json:"rate"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	b.ID = aux.ID
	b.Title = aux.Title
	b.Author = aux.Author
	b.Year = aux.Year
	b.Size = aux.Size
	b.Rate = aux.Rate
	return nil
}

func serializeProto(book *bookpb.Book) ([]byte, error) {
	return proto.Marshal(book)
}

func deserializeProto(data []byte) (*bookpb.Book, error) {
	var b bookpb.Book
	if err := proto.Unmarshal(data, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func serializeBooksToJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func deserializeBooksFromJSON(data []byte) ([]Book, error) {
	var books []Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, err
	}
	return books, nil
}

func serializeBooksToProto(books []*bookpb.Book) ([]byte, error) {
	return proto.Marshal(&bookpb.BookList{Books: books})
}

func deserializeBooksFromProto(data []byte) ([]*bookpb.Book, error) {
	var bookList bookpb.BookList
	if err := proto.Unmarshal(data, &bookList); err != nil {
		return nil, err
	}
	return bookList.Books, nil
}

func main() {
	books := []Book{
		{ID: 1, Title: "Go Programming", Author: "John Doe", Year: 2020, Size: 500.0, Rate: 4.5},
		{ID: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2021, Size: 300.0, Rate: 4.2},
	}

	jsonData, err := serializeBooksToJSON(books)
	if err != nil {
		log.Fatalf("Ошибка сериализации в JSON: %v", err)
	}
	fmt.Println("Сериализация слайса книг в JSON:", string(jsonData))

	deserializedBooks, err := deserializeBooksFromJSON(jsonData)
	if err != nil {
		log.Fatalf("Ошибка десериализации из JSON: %v", err)
	}
	fmt.Println("Десериализация слайса книг из JSON:", deserializedBooks)

	protoBook := &bookpb.Book{
		Id:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2020,
		Size:   500.0,
		Rate:   4.5,
	}

	protoData, err := serializeProto(protoBook)
	if err != nil {
		log.Fatalf("Ошибка сериализации в protobuf: %v", err)
	}
	fmt.Println("Сериализация книги в protobuf:", protoData)

	deserializedProtoBook, err := deserializeProto(protoData)
	if err != nil {
		log.Fatalf("Ошибка десериализации из protobuf: %v", err)
	}
	fmt.Println("Десериализация книги из protobuf:", deserializedProtoBook)

	protoBooks := []*bookpb.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2020, Size: 500.0, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2021, Size: 300.0, Rate: 4.2},
	}

	protoBooksData, err := serializeBooksToProto(protoBooks)
	if err != nil {
		log.Fatalf("Ошибка сериализации слайса книг в protobuf: %v", err)
	}
	fmt.Println("Сериализация слайса книг в protobuf:", protoBooksData)

	deserializedProtoBooks, err := deserializeBooksFromProto(protoBooksData)
	if err != nil {
		log.Fatalf("Ошибка десериализации слайса книг из protobuf: %v", err)
	}
	fmt.Println("Десериализация слайса книг из protobuf:", deserializedProtoBooks)
}
