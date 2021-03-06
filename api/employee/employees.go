package employee

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"launchpad-api/pkg/http_res"
	"launchpad-api/pkg/util"
	"launchpad-api/services"
	"log"
	"net/http"
)

type EmployeeStruct struct {
	Emp_ID     *string `json:"emp_id"`
	Fname      *string `json:"fname"`
	Mname      *string `json:"mname"`
	Lname      *string `json:"lname"`
	Address    *string `json:"address"`
	Dep_ID     *string `json:"dep_id"`
	Created_At *string `json:"created_at"`
	Title      *string `json:"title"`
	Salary     *int    `json:"salary"`
}

//Routes http request to correct method
func HandleAllEmployees(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllEmployees(writer)
	case "PUT":
		res = AddEmployee(writer, req)
	}
	return res
}

//Routes http request to correct method
func HandleEmployee(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetEmployee(writer, req)
	case "POST":
		res = UpdateEmployee(writer, req)
	case "DELETE":
		res = DeleteEmployee(writer, req)
	}
	return res
}

//Get all rows in Employee table
func GetAllEmployees(writer http.ResponseWriter) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Employee")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

//Get a specific row in Employee table
func GetEmployee(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["emp_id"])
	row := services.Db.QueryRow("Select * From Employee Where emp_id = ?", user)
	rowStruct := Row(row)
	if rowStruct.Emp_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No employee exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

//Add a employee to  Employee table
func AddEmployee(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)

	var emp_id string
	for ok := true; ok; ok = util.ValidateEmployeeID(emp_id) {
		emp_id = util.GenerateRandomString(9)
	}

	if _, ok := (*reqMap)["fname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["mname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["lname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["address"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["dep_id"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["title"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["salary"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	fname := (*reqMap)["fname"].(string)
	mname := (*reqMap)["mname"].(string)
	lname := (*reqMap)["lname"].(string)
	address := (*reqMap)["address"].(string)
	dep_id := (*reqMap)["dep_id"].(string)
	title := (*reqMap)["title"].(string)
	salary, _ := (*reqMap)["salary"].(int)

	newEmp := EmployeeStruct{
		Emp_ID:  &emp_id,
		Fname:   &fname,
		Mname:   &mname,
		Lname:   &lname,
		Address: &address,
		Dep_ID:  &dep_id,
		Title:   &title,
		Salary:  &salary,
	}

	if !util.ValidateUpdate(*reqMap, "employee") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Input is invalid"))
	}

	stmt, _ := services.Db.Prepare("Insert into Employee (emp_id,fname,mname,lname,address,dep_id,title,salary) values(?, ?, ?, ?, ?, ?, ?, ?)")
	_, err := stmt.Exec(emp_id, fname, mname, lname, address, dep_id, title, salary)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newEmp)
}

//Take in various attributes and update them in a specific row of  Employee table 
func UpdateEmployee(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "employee") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("employee", (vars["emp_id"]), "emp_id", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}

//Delete a specific row of Employee table
func DeleteEmployee(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["emp_id"])
	row := services.Db.QueryRow("Delete From Employee Where emp_id = ?", user)
	rowStruct := Row(row)
	if rowStruct.Emp_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusOK, errors.New("Employee has been deleted"))
	}
	return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Something went wrong"))
}
