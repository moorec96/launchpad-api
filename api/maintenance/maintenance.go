package maintenance

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

type MaintenanceStruct struct {
	Maint_ID   *string `json:"maint_id"`
	Emp_No     *string `json:"emp_no"`
	Created_at *string `json:"created_at"`
}

func HandleAllMaintenances(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllMaintenance(writer, req)
	case "PUT":
		res = AddMaintenance(writer, req)
	}
	return res
}

func HandleMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetMaintenance(writer, req)
	case "POST":
		res = UpdateMaintenance(writer, req)
	case "DELETE":
		res = DeleteMaintenance(writer, req)
	}
	return res
}

func GetAllMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Maintenance")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

func GetMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	maint := (vars["maint_id"])
	row := services.Db.QueryRow("Select * From Maintenance Where maint_id = ?", maint)
	rowStruct := Row(row)
	if rowStruct.Maint_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No maintenance exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

func AddMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)

	var maint_id string
	for ok := true; ok; ok = util.ValidateMaintenanceID(maint_id) {
		maint_id = util.GenerateRandomString(9)
	}
	if _, ok := (*reqMap)["emp_no"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	emp_no := (*reqMap)["emp_no"].(string)
	newMaint := MaintenanceStruct{
		Maint_ID: &maint_id,
		Emp_No:   &emp_no,
	}
	if !util.ValidateEmployeeID(emp_no) {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No employee with that ID"))
	}

	stmt, _ := services.Db.Prepare("Insert into Maintenance (Maint_ID, EmpNo) values(?, ?)")
	_, err := stmt.Exec(maint_id, emp_no)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Badder Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newMaint)
}

func UpdateMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "maintenance") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("maintenance", (vars["maint_id"]), "maint_id", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Badder Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}

func DeleteMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["maint_id"])
	row := services.Db.QueryRow("Delete From Maintenance Where maint_id = ?", user)
	rowStruct := Row(row)
	if rowStruct.Maint_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusOK, errors.New("Maintenance has been deleted"))
	}
	return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Something went wrong"))
}
