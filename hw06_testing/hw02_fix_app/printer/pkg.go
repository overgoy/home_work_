package printer

import (
	"fmt"
	"github.com/overgoy/home_work_/hw06_testing/hw02/type)

func FormatEmployeeData(employee types.Employee) string {
	return fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
		employee.UserID, employee.Age, employee.Name, employee.DepartmentID)
}

func PrintStaff(staff []types.Employee) {
	for _, employee := range staff {
		str := FormatEmployeeData(employee)
		fmt.Println(str)
	}
}
