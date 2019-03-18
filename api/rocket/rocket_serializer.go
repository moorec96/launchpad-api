package rocket

import (
	"database/sql"
)

//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]RocketStruct {
	output := make([]RocketStruct, 0)
	for rows.Next() {
		var rocketRow RocketStruct
		_ = rows.Scan(&rocketRow.R_ID, &rocketRow.Rname, &rocketRow.Last_Maint, &rocketRow.Created_At)
		output = append(output, rocketRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *RocketStruct {
	var rocketRow RocketStruct
	_ = row.Scan(&rocketRow.R_ID, &rocketRow.Rname, &rocketRow.Last_Maint, &rocketRow.Created_At)
	return &rocketRow
}
