package util

import (
	//"launchpad-api/api/department"
	"launchpad-api/services"
)

var employeeValidationMethods = map[string]interface{}{
	"emp_id":     preventChange,
	"fname":      ValidateString,
	"mname":      ValidateString,
	"lname":      ValidateString,
	"address":    ValidateString,
	"dep_id":     ValidateDepartmentID,
	"created_at": preventChange,
	"title":      ValidateString,
	"salary":     ValidateInt,
}

var departmentValidationMethods = map[string]interface{}{
	"dnum":       preventChange,
	"dname":      ValidateString,
	"created_at": preventChange,
	"head":       ValidateEmployeeID,
}

var inventoryValidationMethods = map[string]interface{}{
	"part_id":    preventChange,
	"quantity":   ValidateInt,
	"created_at": preventChange,
	"location":   ValidateString,
}

var maintenanceValidationMethods = map[string]interface{}{
	"maint_id":   preventChange,
	"emp_no":     ValidateEmployeeID,
	"created_at": preventChange,
}

var rocketValidationMethods = map[string]interface{}{
	"r_id":       preventChange,
	"rname":      ValidateString,
	"last_maint": ValidateMaintenanceID,
	"created_at": preventChange,
}

var rocketLaunchValidationMethods = map[string]interface{}{
	"launch_id":  preventChange,
	"rocket_id":  ValidateRocketID,
	"rlname":     ValidateString,
	"launcher":   ValidateEmployeeID,
	"location":   ValidateString,
	"created_at": preventChange,
}

var partValidationMethods = map[string]interface{}{
	"pnum":  preventChange,
	"pname": ValidateString,
}

//Take in all columns to be updated and ensure updates are good
func ValidateUpdate(newValues map[string]interface{}, table string) bool {
	validationMethods := getValidationMethods(table)
	for key, value := range newValues {
		if _, ok := validationMethods[key]; !ok {
			return false
		}
		if !(validationMethods[key]).(func(interface{}) bool)(value) {
			return false
		}
	}
	return true
}

//String can not be empty, or over 255 in length
func ValidateString(val interface{}) bool {
	valStr := val.(string)
	if len(valStr) < 0 || len(valStr) > 255 {
		return false
	}
	return true
}

//int cannot be below 0
func ValidateInt(val interface{}) bool {
	salary := val.(float64)
	if salary < 0 {
		return false
	}
	return true
}

//Ensure emp_ID exists
func ValidateEmployeeID(val interface{}) bool {
	valStr := val.(string)
	rows, err := services.Db.Query("Select * From Employee Where emp_id = ?", valStr)
	if err != nil {
		return false
	}
	if rows.Next() {
		rows.Close()
		return true
	}
	return false
}

//Ensure Maiint_ID exists
func ValidateMaintenanceID(val interface{}) bool {
	valStr := val.(string)
	rows, err := services.Db.Query("Select * From Maintenance Where maint_id = ?", valStr)
	if err != nil {
		return false
	}
	if rows.Next() {
		rows.Close()
		return true
	}
	return false
}

//Ensure Dep_ID exists
func ValidateDepartmentID(val interface{}) bool {
	valStr := val.(string)
	rows, err := services.Db.Query("Select * From Department Where dnum = ?", valStr)
	if err != nil {
		return false
	}
	if rows.Next() {
		rows.Close()
		return true
	}
	return false
}

//Ensure Pnum exists
func ValidatePartID(val interface{}) bool {
	valStr := val.(string)
	rows, err := services.Db.Query("Select * From Part Where Pnum = ?", valStr)
	if err != nil {
		return false
	}
	if rows.Next() {
		rows.Close()
		return true
	}
	return false
}

//Ensure R_ID Exists
func ValidateRocketID(val interface{}) bool {
	valStr := val.(string)
	rows, err := services.Db.Query("Select * From Rocket Where R_ID = ?", valStr)
	if err != nil {
		return false
	}
	if rows.Next() {
		rows.Close()
		return true
	}
	return false
}

//Ensure Launch_ID exists
func ValidateLaunchID(val interface{}) bool {
	valStr := val.(string)
	rows, err := services.Db.Query("Select * From Rocket_Launch Where Launch_ID = ?", valStr)
	if err != nil {
		return false
	}
	if rows.Next() {
		rows.Close()
		return true
	}
	return false
}

//Any attribute that should not be changed  i.e. IDs 
func preventChange(val interface{}) bool {
	return false
}

//Determines which map of validation methods should be called
func getValidationMethods(table string) map[string]interface{} {
	switch table {
	case "employee":
		return employeeValidationMethods
	case "department":
		return departmentValidationMethods
	case "inventory":
		return inventoryValidationMethods
	case "maintenance":
		return maintenanceValidationMethods
	case "rocket":
		return rocketValidationMethods
	case "rocket_launch":
		return rocketLaunchValidationMethods
	case "part":
		return partValidationMethods
	}
	return nil
}
