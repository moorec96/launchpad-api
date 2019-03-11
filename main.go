package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/{companyId}", HandleCompanyData).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/employees", HandleAllEmployees).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/employees/{employeeId}", HandleEmployee).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/departments/{departmentId}", HandleDepartment).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/departments", HandleAllDepartments).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/maintenance/{maintenanceId}", HandleMaintenanceRecord).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/maintenance", HandleAllMaintenanceRecords).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/rockets/{rocketId}", GetRocket).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/rockets", handleAllRockets).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/inventory/{partId}", GetPart).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/rockets/{rocketId}/parts", handleRocketParts).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/launch", handleAllLaunches).Methods("GET", "PUT")
	router.HandleFunc("/{companyId}/launch/{launchId}", GetLaunch).Methods("GET", "POST")
	router.HandleFunc("/{companyId}/inventory", GetInventory).Methods("GET", "POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
