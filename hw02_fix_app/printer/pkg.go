package printer

import (
	"fmt"
	types "github.com/overgoy/home_work_/hw02_fix_app/types"
)

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
