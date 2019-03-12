package Employee

import (
	"launchpad-api/pkg/http_res"
	"launchpad-api/services"
	"log"
	"net/http"
)

func HandleAllEmployees(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var res *http_res.HttpResponse
	switch req.Method {
	case "GET":
		res = GetAllEmployees(writer, req)
	case "PUT":

	}
	return res
}

func GetAllEmployees(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	rows, err := services.Db.Query("Select fname From Employee")
	var fname string
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&fname)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fname)
	}
	return http_res.GenerateHttpResponse(http.StatusOK, rows)
}
