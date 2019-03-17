package rocket_launch

import (
	"database/sql"
)

func Rows(rows *sql.Rows) *[]RocketLaunchStruct {
	output := make([]RocketLaunchStruct, 0)
	for rows.Next() {
		var launchRow RocketLaunchStruct
		_ = rows.Scan(&launchRow.Launch_ID, &launchRow.Rocket_ID, &launchRow.RLname, &launchRow.Launcher, &launchRow.Location)
		output = append(output, launchRow)
	}
	return &output
}

func Row(row *sql.Row) *RocketLaunchStruct {
	var launchRow RocketLaunchStruct
	_ = row.Scan(&launchRow.Launch_ID, &launchRow.Rocket_ID, &launchRow.RLname, &launchRow.Launcher, &launchRow.Location)
	return &launchRow
}
