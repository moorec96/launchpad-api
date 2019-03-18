package part

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

type PartStruct struct {
	Pnum  *string `json:"pnum"`
	Pname *string `json:"pname"`
}

//Routes http request to correct method
func HandleAllParts(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllParts(writer, req)
	case "PUT":
		res = AddPart(writer, req)
	}
	return res
}

//Routes http request to correct method
func HandlePart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetPart(writer, req)
	case "POST":
		res = UpdatePart(writer, req)
	case "DELETE":
		res = DeletePart(writer, req)
	}
	return res
}

//Get all rows in Parts table
func GetAllParts(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Part")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

func GetPart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	part := (vars["pnum"])
	row := services.Db.QueryRow("Select * From Part Where pnum = ?", part)
	rowStruct := Row(row)
	if rowStruct.Pnum == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No part exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

//Get a specific row in Parts table
func AddPart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {

	reqMap := util.RequestBodyAsMap(req)

	//Loop through random emp_id until one is not taken
	var pnum string
	for ok := true; ok; ok = util.ValidatePartID(pnum) {
		pnum = util.GenerateRandomString(9)
	}
	if _, ok := (*reqMap)["pname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	pname := (*reqMap)["pname"].(string)

	newPart := PartStruct{
		Pnum:  &pnum,
		Pname: &pname,
	}

	if !util.ValidateUpdate(*reqMap, "part") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Input is invalid"))
	}

	stmt, _ := services.Db.Prepare("Insert into Part (pnum, pname) values(?, ?)")
	_, err := stmt.Exec(pnum, pname)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newPart)
}

//Take in various attributes and update them in a specific row of  Parts table 
func UpdatePart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "part") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("part", (vars["pnum"]), "pnum", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Badder Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}

//Delete a specific row of Parts table
func DeletePart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["pnum"])
	row := services.Db.QueryRow("Delete From Part Where pnum = ?", user)
	rowStruct := Row(row)
	if rowStruct.Pnum == nil {
		return http_res.GenerateHttpResponse(http.StatusOK, errors.New("Part has been deleted"))
	}
	return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Something went wrong"))
}
