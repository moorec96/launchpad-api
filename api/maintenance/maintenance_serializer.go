package maintenance

import (
	"database/sql"
)
//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]MaintenanceStruct {
	output := make([]MaintenanceStruct, 0)
	for rows.Next() {
		var maintenanceRow MaintenanceStruct
		_ = rows.Scan(&maintenanceRow.Maint_ID, &maintenanceRow.Emp_No, &maintenanceRow.Created_at)
		output = append(output, maintenanceRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *MaintenanceStruct {
	var maintenanceRow MaintenanceStruct
	_ = row.Scan(&maintenanceRow.Maint_ID, &maintenanceRow.Emp_No, &maintenanceRow.Created_at)
	return &maintenanceRow
}
