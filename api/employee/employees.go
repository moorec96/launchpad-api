package employee

import (
	"fmt"
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
	Created_At *string `json:"date"`
	Title      *string `json:"title"`
	Salary     *int    `json:"salary"`
}

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

func GetAllEmployees(writer http.ResponseWriter) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Employee")
	//var fname string
	if err != nil {
		log.Fatal(err)
	}

	rowStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

func AddEmployee(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	fmt.Print(*reqMap)
	emp_id := "1000"
	fname := (*reqMap)["fname"].(string)
	mname := (*reqMap)["mname"].(string)
	lname := (*reqMap)["lname"].(string)
	address := (*reqMap)["address"].(string)
	dep_id := (*reqMap)["dep_id"].(string)
	created_at := (*reqMap)["created_at"].(string)
	title := (*reqMap)["title"].(string)
	salary, _ := (*reqMap)["salary"].(int)

	newEmp := EmployeeStruct{
		Emp_ID:     &emp_id,
		Fname:      &fname,
		Mname:      &mname,
		Lname:      &lname,
		Address:    &address,
		Dep_ID:     &dep_id,
		Created_At: &created_at,
		Title:      &title,
		Salary:     &salary,
	}

	stmt, _ := services.Db.Prepare("Insert into Employee values(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	_, _ = stmt.Exec(emp_id, fname, mname, lname, address, dep_id, created_at, title, salary)
	return http_res.GenerateHttpResponse(http.StatusOK, newEmp)
}
