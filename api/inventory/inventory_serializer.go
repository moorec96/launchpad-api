package inventory

import (
	"database/sql"
)

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

func Row(row *sql.Row) *InventoryStruct {
	var inventoryRow InventoryStruct
	_ = row.Scan(&inventoryRow.Part_ID, &inventoryRow.Quantity,
		&inventoryRow.Created_at, &inventoryRow.Location)
	return &inventoryRow
}
