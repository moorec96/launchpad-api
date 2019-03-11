package Company

import (
	"fmt"
	"github.com/gorilla/mux"
	"launchpad-api/pkg/http_res"
	"net/http"
)

func HandleCompanyData(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {
	var resp = *http_res.HttpResponse

	switch req.Method {
	case "GET":
		res = GetCompanyData(writer, req)
	case "POST":

	}
	return resp
}

func GetCompanyData(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse {

}
