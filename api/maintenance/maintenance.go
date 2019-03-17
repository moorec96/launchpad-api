package maintenance

import (
	"errors"
	"launchpad-api/pkg/http_res"
	"launchpad-api/services"
	"log"
	"net/http"
)

type MaintenanceStruct struct {
	Maint_ID   *string `json:"maint_id"`
	Emp_No     *string `json:"emp_no"`
	Created_at *string `json:"created_at"`
}

func HandleMaintenance(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllMaintenance(writer, req)
	case "PUT":
		//res = AddMaintenance(writer, req)
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
