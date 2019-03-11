package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleRoutes() {
	router := mux.NewRouter()

	//Company
	router.HandleFunc("/{companyId}", HandleCompanyData).Methods("GET", "POST")

	//Employees
	router.HandleFunc("/{companyId}/employees", HandleAllEmployees).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/employees/{employeeId}", HandleEmployee).Methods("GET", "POST")

	//Departments
	router.HandleFunc("/{companyId}/departments", HandleAllDepartments).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/departments/{departmentId}", HandleDepartment).Methods("GET", "POST")

	//Maintenance
	router.HandleFunc("/{companyId}/maintenance", HandleAllMaintenanceRecords).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/maintenance/{maintenanceId}", HandleMaintenanceRecord).Methods("GET", "POST")

	//Rockets
	router.HandleFunc("/{companyId}/rockets", HandleAllRockets).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/rockets/{rocketId}", HandleRocket).Methods("GET", "POST")

	//Launches
	router.HandleFunc("/{companyId}/launch", HandleAllLaunches).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/launch/{launchId}", HandleLaunch).Methods("GET", "POST")

	//Inventory
	router.HandleFunc("/{companyId}/inventory", HandleAllInventory).Methods("GET", "POST")

	//Parts
	router.HandleFunc("/{companyId}/inventory/{partId}", HandlePart).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/rockets/{rocketId}/parts", HandleRocketParts).Methods("GET", "PUT")

	http.Handle("/", r)

}
