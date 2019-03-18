package rocket

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

type RocketStruct struct {
	R_ID       *string `json:"r_id"`
	Rname      *string `json:"rname"`
	Last_Maint *string `json:"last_maint"`
	Created_At *string `json:"created_at"`
}

func HandleAllRockets(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllRockets(writer, req)
	case "PUT":
		res = AddRocket(writer, req)
	}
	return res
}

func HandleRocket(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetRocket(writer, req)
	case "POST":
		res = UpdateRocket(writer, req)
	case "DELETE":
		res = DeleteRocket(writer, req)
	}
	return res
}

func GetAllRockets(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Rocket")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

func GetRocket(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	rocket := (vars["r_id"])
	row := services.Db.QueryRow("Select * From Rocket Where r_id = ?", rocket)
	rowStruct := Row(row)
	if rowStruct.R_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No rocket exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

func AddRocket(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {

	reqMap := util.RequestBodyAsMap(req)

	//Loop through random emp_id until one is not taken
	var r_id string
	for ok := true; ok; ok = util.ValidateRocketID(r_id) {
		r_id = util.GenerateRandomString(9)
	}
	if _, ok := (*reqMap)["rname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	rname := (*reqMap)["rname"].(string)
	newRocket := RocketStruct{
		R_ID:       &r_id,
		Rname:      &rname,
		Last_Maint: nil,
	}

	if !util.ValidateUpdate(*reqMap, "rocket") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Input is invalid"))
	}

	stmt, _ := services.Db.Prepare("Insert into Rocket (r_id, rname) values(?, ?)")
	_, err := stmt.Exec(r_id, rname)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newRocket)
}

func UpdateRocket(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "rocket") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("rocket", (vars["r_id"]), "r_id", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Badder Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}

func DeleteRocket(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["r_id"])
	row := services.Db.QueryRow("Delete From Rocket Where r_id = ?", user)
	rowStruct := Row(row)
	if rowStruct.R_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusOK, errors.New("Rocket has been deleted"))
	}
	return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Something went wrong"))
}
