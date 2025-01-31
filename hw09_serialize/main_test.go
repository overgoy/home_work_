package main

import (
	"testing"

	"github.com/overgoy/home_work_/hw09_serialize/bookpb" // Путь импорта для сгенерированного protobuf пакета
)

func TestSerializeBooksToJSON(t *testing.T) {
	books := []Book{
		{ID: 1, Title: "Go Programming", Author: "John Doe", Year: 2020, Size: 500.0, Rate: 4.5},
		{ID: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2021, Size: 300.0, Rate: 4.2},
	}

	data, err := serializeBooksToJSON(books)
	if err != nil {
		t.Fatalf("Ошибка сериализации в JSON: %v", err)
	}

	if len(data) == 0 {
		t.Error("Сериализованный JSON не должен быть пустым")
	}
}

func TestDeserializeBooksFromJSON(t *testing.T) {
	books := []Book{
		{ID: 1, Title: "Go Programming", Author: "John Doe", Year: 2020, Size: 500.0, Rate: 4.5},
		{ID: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2021, Size: 300.0, Rate: 4.2},
	}

	data, err := serializeBooksToJSON(books)
	if err != nil {
		t.Fatalf("Ошибка сериализации в JSON: %v", err)
	}

	deserializedBooks, err := deserializeBooksFromJSON(data)
	if err != nil {
		t.Fatalf("Ошибка десериализации из JSON: %v", err)
	}

	if len(deserializedBooks) != len(books) {
		t.Errorf("Ожидалось %d книг, но получено %d", len(books), len(deserializedBooks))
	}
}

func TestSerializeProto(t *testing.T) {
	protoBook := &bookpb.Book{
		Id:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2020,
		Size:   500.0,
		Rate:   4.5,
	}

	data, err := serializeProto(protoBook)
	if err != nil {
		t.Fatalf("Ошибка сериализации в protobuf: %v", err)
	}

	if len(data) == 0 {
		t.Error("Сериализованные данные protobuf не должны быть пустыми")
	}
}

func TestDeserializeProto(t *testing.T) {
	protoBook := &bookpb.Book{
		Id:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2020,
		Size:   500.0,
		Rate:   4.5,
	}

	data, err := serializeProto(protoBook)
	if err != nil {
		t.Fatalf("Ошибка сериализации в protobuf: %v", err)
	}

	deserializedProtoBook, err := deserializeProto(data)
	if err != nil {
		t.Fatalf("Ошибка десериализации из protobuf: %v", err)
	}

	if deserializedProtoBook.GetId() != protoBook.GetId() {
		t.Errorf("Ожидалось ID %d, но получено %d", protoBook.GetId(), deserializedProtoBook.GetId())
	}
}

func TestSerializeBooksToProto(t *testing.T) {
	protoBooks := []*bookpb.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2020, Size: 500.0, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2021, Size: 300.0, Rate: 4.2},
	}

	data, err := serializeBooksToProto(protoBooks)
	if err != nil {
		t.Fatalf("Ошибка сериализации слайса книг в protobuf: %v", err)
	}

	if len(data) == 0 {
		t.Error("Сериализованные данные protobuf для слайса книг не должны быть пустыми")
	}
}

func TestDeserializeBooksFromProto(t *testing.T) {
	protoBooks := []*bookpb.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2020, Size: 500.0, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2021, Size: 300.0, Rate: 4.2},
	}

	data, err := serializeBooksToProto(protoBooks)
	if err != nil {
		t.Fatalf("Ошибка сериализации слайса книг в protobuf: %v", err)
	}

	deserializedProtoBooks, err := deserializeBooksFromProto(data)
	if err != nil {
		t.Fatalf("Ошибка десериализации слайса книг из protobuf: %v", err)
	}

	if len(deserializedProtoBooks) != len(protoBooks) {
		t.Errorf("Ожидалось %d книг, но получено %d", len(protoBooks), len(deserializedProtoBooks))
	}
}
