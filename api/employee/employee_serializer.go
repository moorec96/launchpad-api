package employee

import (
	"database/sql"
	"fmt"
)

func Rows(rows *sql.Rows) *[]EmployeeStruct {
	output := make([]EmployeeStruct, 0)
	for rows.Next() {
		var employeeRow EmployeeStruct
		_ = rows.Scan(&employeeRow.Emp_ID, &employeeRow.Fname,
			&employeeRow.Lname, &employeeRow.Address, &employeeRow.Dep_ID,
			&employeeRow.Created_At, &employeeRow.Title, &employeeRow.Salary)
		fmt.Print(employeeRow)
		output = append(output, employeeRow)
	}
	return &output
}
