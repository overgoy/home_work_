package main

import (
	"github.com/overgoy/home_work_/hw09_serialize/bookpb" // Путь импорта для сгенерированного protobuf пакета
	_ "google.golang.org/protobuf/proto"
	"testing"
)

func TestSerializationJSON(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2020,
		Size:   500.0,
		Rate:   4.5,
	}

	data, err := book.MarshalJSON()
	if err != nil {
		t.Errorf("Ошибка сериализации JSON: %v", err)
	}

	var deserializedBook Book
	err = deserializedBook.UnmarshalJSON(data)
	if err != nil {
		t.Errorf("Ошибка десериализации JSON: %v", err)
	}

	if deserializedBook != book {
		t.Errorf("Ожидалось %v, но получено %v", book, deserializedBook)
	}
}

func TestSerializationProto(t *testing.T) {
	book := &bookpb.Book{
		Id:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2020,
		Size:   500.0,
		Rate:   4.5,
	}

	data, err := serializeProto(book)
	if err != nil {
		t.Errorf("Ошибка сериализации Proto: %v", err)
	}

	deserializedBook, err := deserializeProto(data)
	if err != nil {
		t.Errorf("Ошибка десериализации Proto: %v", err)
	}

	if deserializedBook.GetId() != book.GetId() {
		t.Errorf("Ожидалось %v, но получено %v", book.GetId(), deserializedBook.GetId())
	}

	if deserializedBook.GetTitle() != book.GetTitle() {
		t.Errorf("Ожидалось %v, но получено %v", book.GetTitle(), deserializedBook.GetTitle())
	}
	if deserializedBook.GetAuthor() != book.GetAuthor() {
		t.Errorf("Ожидалось %v, но получено %v", book.GetAuthor(), deserializedBook.GetAuthor())
	}
	if deserializedBook.GetYear() != book.GetYear() {
		t.Errorf("Ожидалось %v, но получено %v", book.GetYear(), deserializedBook.GetYear())
	}
	if deserializedBook.GetSize() != book.GetSize() {
		t.Errorf("Ожидалось %v, но получено %v", book.GetSize(), deserializedBook.GetSize())
	}
	if deserializedBook.GetRate() != book.GetRate() {
		t.Errorf("Ожидалось %v, но получено %v", book.GetRate(), deserializedBook.GetRate())
	}
}
