package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/{companyId}",GetCompanyData).Methods("GET")
	router.HandleFunc("/{companyId}/employees", GetAllEmployees).Methods("GET")
	router.HandleFunc("/{companyId}/employees/{employeeId}", GetEmployee).Methods("GET")
	router.HandleFunc("/{companyId}/departments", GetAllDepartments).Methods("GET")
	router.HandleFunc("/{companyId}/departments/{departmentId}", GetDepartment).Methods("GET")
	router.HandleFunc("/{companyId}/maintenance",GetAllMaintenanceRecords).Methods("GET")
	router.HandleFunc("/{companyId}/maintenance/{maintenanceId}", GetMaintenanceRecord).Methods("GET")
	router.HandleFunc("/{companyId}/rockets", GetAllRockets).Methods("GET")
	router.HandleFunc("/{companyId}/rockets/{rocketId}", GetRocket).Methods("GET")
	router.HandleFunc("/{companyId}/rockets/{rocketId}/parts", GetRocketParts).Methods("GET")
	router.HandleFunc("/{companyId}/launch", GetAllLaunches).Methods("GET")
	router.HandleFunc("/{companyId}/launch/{launchId}", GetLaunch).Methods("GET")
	router.HandleFunc("/{companyId}/inventory", GetInventory).Methods("GET")
	router.HandleFunc("/{companyId}/inventory/{partId}",GetPart).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000",router))
}