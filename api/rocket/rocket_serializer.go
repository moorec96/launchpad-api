package rocket

import (
	"database/sql"
)

func Rows(rows *sql.Rows) *[]RocketStruct {
	output := make([]RocketStruct, 0)
	for rows.Next() {
		var rocketRow RocketStruct
		_ = rows.Scan(&rocketRow.R_ID, &rocketRow.Rname, &rocketRow.Last_Maint, &rocketRow.Created_At)
		output = append(output, rocketRow)
	}
	return &output
}

func Row(row *sql.Row) *RocketStruct {
	var rocketRow RocketStruct
	_ = row.Scan(&rocketRow.R_ID, &rocketRow.Rname, &rocketRow.Last_Maint, &rocketRow.Created_At)
	return &rocketRow
}
