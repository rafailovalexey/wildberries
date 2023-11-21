package api

import "net/http"

type InterfaceEmployeeApi interface {
	GetEmployeeById(http.ResponseWriter, *http.Request)
	CreateEmployee(http.ResponseWriter, *http.Request)

	NotFound(http.ResponseWriter, *http.Request)
	MethodNotAllowed(http.ResponseWriter, *http.Request)
}
