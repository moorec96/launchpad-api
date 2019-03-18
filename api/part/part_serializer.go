package part

import (
	"database/sql"
)
//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]PartStruct {
	output := make([]PartStruct, 0)
	for rows.Next() {
		var partRow PartStruct
		_ = rows.Scan(&partRow.Pnum, &partRow.Pname)
		output = append(output, partRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *PartStruct {
	var partRow PartStruct
	_ = row.Scan(&partRow.Pnum, &partRow.Pname)
	return &partRow
}
