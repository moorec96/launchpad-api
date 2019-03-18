package util

import (
	"fmt"
	"launchpad-api/services"
	"log"
	"strings"
)

func ToJson(mp map[string]interface{}) string {
	//keyList, valueList := getUpdateSets(mp)
	json := "{"
	for key, value := range mp {
		//valString := values.(string)
		pair := fmt.Sprintf("%v:%v", key, value)
		json += pair
		json += ","
	}
	if len(json) > 0 {
		res := json[:len(json)-1]
		res += "}"
		return res
	}
	json += "}"
	return json
}

//Takes in map of values to update, and makes sql update query
func UpdateTable(table string, id string, idType string, newValues map[string]interface{}) error {
	updateSets, updateValues := getUpdateSets(newValues)
	updateValues = append(updateValues, id)
	updateSetsString := strings.Join(updateSets, " = ?, ") + " = ?"
	updateSql := "UPDATE %s SET %s WHERE %s = ?"
	updateQuery := fmt.Sprintf(updateSql, table, updateSetsString, idType)
	fmt.Println(updateQuery)
	stmt, err := services.Db.Prepare(updateQuery)
	if err != nil {
		log.Print(err)
		return err
	}
	_, err = stmt.Exec(updateValues...)
	return err
}

func getUpdateSets(newValues map[string]interface{}) ([]string, []interface{}) {
	var updateSets []string
	var updateValues []interface{}
	for key, value := range newValues {
		updateSets = append(updateSets, key)
		updateValues = append(updateValues, value)
	}
	return updateSets, updateValues
}
