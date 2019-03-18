package employee

import (
	"database/sql"
)

//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]EmployeeStruct {
	output := make([]EmployeeStruct, 0)
	for rows.Next() {
		var employeeRow EmployeeStruct
		_ = rows.Scan(&employeeRow.Emp_ID, &employeeRow.Fname, &employeeRow.Mname,
			&employeeRow.Lname, &employeeRow.Address, &employeeRow.Dep_ID,
			&employeeRow.Created_At, &employeeRow.Title, &employeeRow.Salary)
		output = append(output, employeeRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *EmployeeStruct {
	var employeeRow EmployeeStruct
	_ = row.Scan(&employeeRow.Emp_ID, &employeeRow.Fname, &employeeRow.Mname,
		&employeeRow.Lname, &employeeRow.Address, &employeeRow.Dep_ID,
		&employeeRow.Created_At, &employeeRow.Title, &employeeRow.Salary)
	return &employeeRow
}
