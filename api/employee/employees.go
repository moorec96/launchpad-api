package employee

import (
	"launchpad-api/pkg/http_res"
	"launchpad-api/services"
	"log"
	"net/http"
)

type EmployeeStruct struct {
	Emp_ID     *int    `json:"emp_id"`
	Fname      *string `json:"fname"`
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
