package api

import (
	"github.com/gorilla/mux"
	"launchpad-api/api/department"
	"launchpad-api/api/employee"
	"launchpad-api/api/inventory"
	"launchpad-api/api/maintenance"
	"launchpad-api/api/part"
	"launchpad-api/api/rocket"
	"launchpad-api/api/rocket_launch"
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
	{"/employees/{emp_id}", employee.HandleEmployee, []string{"GET", "POST", "DELETE"}},
	{"/departments", department.HandleAllDepartments, []string{"GET", "PUT"}},
	{"/departments/{dnum}", department.HandleDepartment, []string{"GET", "POST", "DELETE"}},
	{"/inventory", inventory.HandleAllInventory, []string{"GET", "PUT"}},
	{"/inventory/{part_id}", inventory.HandleInventoryPart, []string{"GET", "POST", "DELETE"}},
	{"/maintenance", maintenance.HandleAllMaintenances, []string{"GET", "PUT"}},
	{"/maintenance/{maint_id}", maintenance.HandleMaintenance, []string{"GET", "POST", "DELETE"}},
	{"/rocket", rocket.HandleAllRockets, []string{"GET", "PUT"}},
	{"/rocket/{r_id}", rocket.HandleRocket, []string{"GET", "POST", "DELETE"}},
	{"/rocket_launch", rocket_launch.HandleAllRocketLaunches, []string{"GET", "PUT"}},
	{"/rocket_launch/{launch_id}", rocket_launch.HandleRocketLaunch, []string{"GET", "POST", "DELETE"}},
	{"/part", part.HandleAllParts, []string{"GET", "PUT"}},
	{"/part/{pnum}", part.HandlePart, []string{"GET", "POST", "DELETE"}},
}

//Take in all routes and map them to their specific endpoint functions
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
