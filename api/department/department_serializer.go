package department

import (
	"database/sql"
)

func Rows(rows *sql.Rows) *[]DepartmentStruct {
	output := make([]DepartmentStruct, 0)
	for rows.Next() {
		var departmentRow DepartmentStruct
		_ = rows.Scan(&departmentRow.Dnum, &departmentRow.Dname, &departmentRow.Created_At, &departmentRow.Head)
		output = append(output, departmentRow)
	}
	return &output
}

func Row(row *sql.Row) *DepartmentStruct {
	var departmentRow DepartmentStruct
	_ = row.Scan(&departmentRow.Dnum, &departmentRow.Dname, &departmentRow.Created_At, &departmentRow.Head)
	return &departmentRow
}
