package rocket_launch

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

type RocketLaunchStruct struct {
	Launch_ID  *string `json:"launch_id"`
	Rocket_ID  *string `json:"rocket_id"`
	RLname     *string `json:"rlname"`
	Launcher   *string `json:"launcher"`
	Location   *string `json:"location"`
	Created_At *string `json:"created_at"`
}

func HandleAllRocketLaunches(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllRocketLaunches(writer, req)
	case "PUT":
		res = AddRocketLaunch(writer, req)
	}
	return res
}

func HandleRocketLaunch(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetRocketLaunch(writer, req)
	case "POST":
		res = UpdateLaunch(writer, req)
	case "DELETE":
		res = DeleteRocketLaunch(writer, req)
	}
	return res
}

func GetAllRocketLaunches(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select * From Rocket_Launch")
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad query"))
	}

	rowsStruct := Rows(rows)
	return http_res.GenerateHttpResponse(http.StatusOK, *rowsStruct)
}

func GetRocketLaunch(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	rocket := (vars["launch_id"])
	row := services.Db.QueryRow("Select * From Rocket_Launch Where launch_id = ?", rocket)
	rowStruct := Row(row)
	if rowStruct.Launch_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("No launch exists with that ID"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, *rowStruct)
}

func AddRocketLaunch(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {

	reqMap := util.RequestBodyAsMap(req)

	var launch_id string
	for ok := true; ok; ok = util.ValidateLaunchID(launch_id) {
		launch_id = util.GenerateRandomString(9)
	}

	if _, ok := (*reqMap)["rocket_id"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input (rocket_id)"))
	}
	if _, ok := (*reqMap)["rlname"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input (rlname)"))
	}
	if _, ok := (*reqMap)["launcher"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input (launcher)"))
	}
	if _, ok := (*reqMap)["location"]; !ok {
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input (location)"))
	}
	rocket_id := (*reqMap)["rocket_id"].(string)
	rlname := (*reqMap)["rlname"].(string)
	launcher := (*reqMap)["launcher"].(string)
	location := (*reqMap)["location"].(string)
	newLaunch := RocketLaunchStruct{
		Launch_ID: &launch_id,
		Rocket_ID: &rocket_id,
		RLname:    &rlname,
		Launcher:  &launcher,
		Location:  &location,
	}

	if !util.ValidateUpdate(*reqMap, "rocket_launch") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Input is invalid"))
	}

	stmt, _ := services.Db.Prepare("Insert into Rocket_Launch (launch_id, rocket_id, rlname, launcher, location) values(?, ?, ?, ?, ?)")
	_, err := stmt.Exec(launch_id, rocket_id, rlname, launcher, location)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, newLaunch)
}

func UpdateLaunch(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	reqMap := util.RequestBodyAsMap(req)
	if !util.ValidateUpdate(*reqMap, "rocket_launch") {
		fmt.Println("Bad input")
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Bad Input"))
	}

	vars := mux.Vars(req)
	err := util.UpdateTable("rocket_launch", (vars["launch_id"]), "launch_id", *reqMap)
	if err != nil {
		log.Print(err)
		return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Badder Input"))
	}
	return http_res.GenerateHttpResponse(http.StatusOK, "Successful Update")
}

func DeleteRocketLaunch(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	vars := mux.Vars(req)
	user := (vars["launch_id"])
	row := services.Db.QueryRow("Delete From Rocket_Launch Where launch_id = ?", user)
	rowStruct := Row(row)
	if rowStruct.Launch_ID == nil {
		return http_res.GenerateHttpResponse(http.StatusOK, errors.New("Launch has been deleted"))
	}
	return http_res.GenerateHttpResponse(http.StatusBadRequest, errors.New("Something went wrong"))
}
