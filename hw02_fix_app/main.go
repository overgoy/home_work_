package main

import (
	"fmt"
	"github.com/fixme_my_friend/hw02_fix_app/types"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

func main() {
	path := "data.json"

	// Запрашиваем путь к файлу
	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json" // используем значение по умолчанию
	}

	// Читаем данные из файла
	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	// Печатаем информацию о сотрудниках
	printer.PrintStaff(staff)
}
