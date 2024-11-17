package hw02

import (
	"fmt"

	"github.com/overgoy/home_work_/hw06_testing/hw02/printer"
	"github.com/overgoy/home_work_/hw06_testing/hw02/reader"
	"github.com/overgoy/home_work_/hw06_testing/hw02/types"
)

func getFilePath() string {
	var path string
	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)
	if len(path) == 0 {
		path = "data.json" // используем значение по умолчанию
	}
	return path
}

func loadStaff(path string) ([]types.Employee, error) {
	return reader.ReadJSON(path)
}

func printStaff(staff []types.Employee) {
	printer.PrintStaff(staff)
}

func main() {
	path := getFilePath()

	staff, err := loadStaff(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	printStaff(staff)
}
