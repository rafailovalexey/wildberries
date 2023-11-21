package api

import "net/http"

type InterfaceEmployeeApi interface {
	EmployeesHandler(http.ResponseWriter, *http.Request)

	GetEmployeeById(http.ResponseWriter, *http.Request)
	CreateEmployee(http.ResponseWriter, *http.Request)
}
