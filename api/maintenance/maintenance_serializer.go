package maintenance

import (
	"database/sql"
)

func Rows(rows *sql.Rows) *[]MaintenanceStruct {
	output := make([]MaintenanceStruct, 0)
	for rows.Next() {
		var maintenanceRow MaintenanceStruct
		_ = rows.Scan(&maintenanceRow.Maint_ID, &maintenanceRow.Emp_No, &maintenanceRow.Created_at)
		output = append(output, maintenanceRow)
	}
	return &output
}

func Row(row *sql.Row) *MaintenanceStruct {
	var maintenanceRow MaintenanceStruct
	_ = row.Scan(&maintenanceRow.Maint_ID, &maintenanceRow.Emp_No, &maintenanceRow.Created_at)
	return &maintenanceRow
}
