package util

import (
	//"launchpad-api/api/department"
	"launchpad-api/services"
)

var employeeValidationMethods = map[string]interface{}{
	"emp_id":     preventChange,
	"fname":      validateString,
	"mname":      validateString,
	"lname":      validateString,
	"address":    validateString,
	"dep_id":     ValidateDepartmentID,
	"created_at": preventChange,
	"title":      validateString,
	"salary":     validateSalary,
}

var departmentValidationMethods = map[string]interface{}{
	"dnum":       preventChange,
	"dname":      validateString,
	"created_at": preventChange,
	"head":       ValidateEmployeeID,
}

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

func validateString(val interface{}) bool {
	valStr := val.(string)
	if len(valStr) < 0 || len(valStr) > 255 {
		return false
	}
	return true
}

func validateSalary(val interface{}) bool {
	salary := val.(float64)
	if salary < 0 {
		return false
	}
	return true
}

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

func preventChange(val interface{}) bool {
	return false
}

func getValidationMethods(table string) map[string]interface{} {
	switch table {
	case "employee":
		return employeeValidationMethods
	case "department":
		return departmentValidationMethods

	}
	return nil
}
