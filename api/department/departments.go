package department

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

type DepartmentStruct struct {
	Dnum       *string `json:"dnum"`
	Dname      *string `json:"dname"`
	Created_At *string `json:"created_at"`
	Head       *string `json:"head"`
}

func HandleAllDepartments(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetDepartments(writer)
	case "PUT":
		res = AddDepartment(writer, req)
	}
	return res
}

func HandleDepartment(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetDepartment(writer, req)
	case "POST":
		res = UpdateDepartment(writer, req)
	}
	return res
}

func GetDepartments(writer http.ResponseWriter) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Department")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

func GetDepartment(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["dnum"])
	row := services.Db.QueryRow("Select * From Department Where dnum = ?", user)
	rowStruct := Row(row)
	if rowStruct.Dnum == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No department exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

func AddDepartment(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if _, ok := (*reqMap)["dnum"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["dname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["head"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	dnum := (*reqMap)["dnum"].(string)
	dname := (*reqMap)["dname"].(string)
	head := (*reqMap)["head"].(string)
	if util.ValidateDepartmentID(dnum) {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("That department num is already taken"))
	}
	if !util.ValidateString(dname) {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if !util.ValidateEmployeeID(head) {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("That head emp_id does not exist"))
	}
	newDep := DepartmentStruct{
		Dnum:  &dnum,
		Dname: &dname,
		Head:  &head,
	}

	stmt, _ := services.Db.Prepare("Insert into Department (dnum,dname,head) values(?, ?, ?)")
	_, err := stmt.Exec(dnum, dname, head)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newDep)
}

func UpdateDepartment(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "department") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Input is invalid"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("department", (vars["dnum"]), "dnum", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}
