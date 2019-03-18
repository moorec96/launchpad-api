package inventory

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

type InventoryStruct struct {
	Part_ID    *string  `json:"part_id"`
	Quantity   *float64 `json:"quantity"`
	Created_at *string  `json:"created_at"`
	Location   *string  `json:"location"`
}

func HandleAllInventory(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllInventory(writer)
	case "PUT":
		res = AddPart(writer, req)
	}
	return res
}

func HandleInventoryPart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetInventoryPart(writer, req)
	case "POST":
		res = UpdatePart(writer, req)
	case "DELETE":
		res = DeleteInventoryItem(writer, req)
	}
	return res
}

func GetAllInventory(writer http.ResponseWriter) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Inventory")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

func GetInventoryPart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	part := (vars["part_id"])
	row := services.Db.QueryRow("Select * From Inventory Where part_id = ?", part)
	rowStruct := Row(row)
	if rowStruct.Part_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No part exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

func AddPart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)

	if _, ok := (*reqMap)["part_id"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["quantity"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	if _, ok := (*reqMap)["location"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	part_id := (*reqMap)["part_id"].(string)
	quantity := (*reqMap)["quantity"].(float64)
	location := (*reqMap)["location"].(string)
	fmt.Print(part_id)
	if !util.ValidatePartID(part_id) {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Part does not exist"))
	}
	if !util.ValidateInt(quantity) {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Quantity input was bad"))
	}
	if !util.ValidateString(location) {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Location input was bad"))
	}

	newInv := InventoryStruct{
		Part_ID:  &part_id,
		Quantity: &quantity,
		Location: &location,
	}

	stmt, _ := services.Db.Prepare("Insert into Inventory (part_id,quantity,location) values(?, ?, ?)")
	_, err := stmt.Exec(part_id, quantity, location)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newInv)
}

func UpdatePart(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "inventory") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("inventory", (vars["part_id"]), "part_id", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Badder Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}

func DeleteInventoryItem(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["part_id"])
	row := services.Db.QueryRow("Delete From Inventory Where part_id = ?", user)
	rowStruct := Row(row)
	if rowStruct.Part_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusOK, errors.New("Part has been deleted"))
	}
	return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Something went wrong"))
}
