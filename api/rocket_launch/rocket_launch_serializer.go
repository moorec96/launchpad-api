package rocket_launch

import (
	"database/sql"
)

//Takes in return value of a sql query, and places it into an array of structs
func Rows(rows *sql.Rows) *[]RocketLaunchStruct {
	output := make([]RocketLaunchStruct, 0)
	for rows.Next() {
		var launchRow RocketLaunchStruct
		_ = rows.Scan(&launchRow.Launch_ID, &launchRow.Rocket_ID, &launchRow.RLname, &launchRow.Launcher, &launchRow.Location, &launchRow.Created_At)
		output = append(output, launchRow)
	}
	return &output
}

//Takes in return value of a sql query, and places it into an array of structs. Used for single row queries
func Row(row *sql.Row) *RocketLaunchStruct {
	var launchRow RocketLaunchStruct
	_ = row.Scan(&launchRow.Launch_ID, &launchRow.Rocket_ID, &launchRow.RLname, &launchRow.Launcher, &launchRow.Location, &launchRow.Created_At)
	return &launchRow
}
