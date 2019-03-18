package inventory

import (
	"database/sql"
)

//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]InventoryStruct {
	output := make([]InventoryStruct, 0)
	for rows.Next() {
		var inventoryRow InventoryStruct
		_ = rows.Scan(&inventoryRow.Part_ID, &inventoryRow.Quantity, &inventoryRow.Created_at,
			&inventoryRow.Location)
		output = append(output, inventoryRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *InventoryStruct {
	var inventoryRow InventoryStruct
	_ = row.Scan(&inventoryRow.Part_ID, &inventoryRow.Quantity,
		&inventoryRow.Created_at, &inventoryRow.Location)
	return &inventoryRow
}
