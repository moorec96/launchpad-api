package department

import (
	"database/sql"
)
//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]DepartmentStruct {
	output := make([]DepartmentStruct, 0)
	for rows.Next() {
		var departmentRow DepartmentStruct
		_ = rows.Scan(&departmentRow.Dnum, &departmentRow.Dname, &departmentRow.Created_At, &departmentRow.Head)
		output = append(output, departmentRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *DepartmentStruct {
	var departmentRow DepartmentStruct
	_ = row.Scan(&departmentRow.Dnum, &departmentRow.Dname, &departmentRow.Created_At, &departmentRow.Head)
	return &departmentRow
}
