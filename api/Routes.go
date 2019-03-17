package api

import (
	"github.com/gorilla/mux"
	"launchpad-api/api/department"
	"launchpad-api/api/employee"
	"launchpad-api/pkg/http_res"
	"net/http"
)

type Endpoint struct {
	Path     string
	Function func(writer http.ResponseWriter, req *http.Request) *http_res.HttpResponse
	Methods  []string
}

var endpoints = []Endpoint{
	{"/employees", employee.HandleAllEmployees, []string{"GET", "PUT"}},
	{"/employees/{emp_id}", employee.HandleEmployee, []string{"GET", "POST"}},
	{"/departments", department.HandleAllDepartments, []string{"GET", "PUT"}},
	{"/departments/{dnum}", department.HandleDepartment, []string{"GET", "POST"}},
}

func HandleRoutes() {
	r := mux.NewRouter()

	for _, endpoint := range endpoints {
		handleEndpoint(r, endpoint)
	}

	http.Handle("/", r)
}

func handleEndpoint(r *mux.Router, endpoint Endpoint) {
	r.HandleFunc(endpoint.Path, func(writer http.ResponseWriter, req *http.Request) {
		mapEndpoints(writer, req, endpoint)
	}).Methods(endpoint.Methods...)
}

func mapEndpoints(writer http.ResponseWriter, req *http.Request, endpoint Endpoint) {
	res := endpoint.Function(writer, req)
	if res == nil {
		res = http_res.GenerateHttpResponse(http.StatusOK, "Ok")
	}
	http_res.SendHttpResponse(writer, res)
}

//func HandleRoutes() {
//	router := mux.NewRouter()

//Departments
//router.HandleFunc("/{companyId}/departments", HandleAllDepartments).Methods("GET", "PUT")
//router.HandleFunc("/{companyId}/departments/{departmentId}", HandleDepartment).Methods("GET", "POST")

//Maintenance
//router.HandleFunc("/{companyId}/maintenance", HandleAllMaintenanceRecords).Methods("GET", "PUT")
//router.HandleFunc("/{companyId}/maintenance/{maintenanceId}", HandleMaintenanceRecord).Methods("GET", "POST")

//Rockets
//router.HandleFunc("/{companyId}/rockets", HandleAllRockets).Methods("GET", "PUT")
//router.HandleFunc("/{companyId}/rockets/{rocketId}", HandleRocket).Methods("GET", "POST")

//Launches
//router.HandleFunc("/{companyId}/launch", HandleAllLaunches).Methods("GET", "PUT")
//router.HandleFunc("/{companyId}/launch/{launchId}", HandleLaunch).Methods("GET", "POST")

//Inventory
//router.HandleFunc("/{companyId}/inventory", HandleAllInventory).Methods("GET", "POST")

//Parts
//router.HandleFunc("/{companyId}/inventory/{partId}", HandlePart).Methods("GET", "POST")
//router.HandleFunc("/{companyId}/rockets/{rocketId}/parts", HandleRocketParts).Methods("GET", "PUT")

//	http.Handle("/", router)

//}
