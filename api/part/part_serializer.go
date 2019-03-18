package part

import (
	"database/sql"
)

func Rows(rows *sql.Rows) *[]PartStruct {
	output := make([]PartStruct, 0)
	for rows.Next() {
		var partRow PartStruct
		_ = rows.Scan(&partRow.Pnum, &partRow.Pname)
		output = append(output, partRow)
	}
	return &output
}

func Row(row *sql.Row) *PartStruct {
	var partRow PartStruct
	_ = row.Scan(&partRow.Pnum, &partRow.Pname)
	return &partRow
}
