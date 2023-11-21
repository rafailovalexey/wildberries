package employees

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/employees_proxy/internal/api"
	"github.com/emptyhopes/employees_proxy/internal/converter"
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	"github.com/emptyhopes/employees_proxy/internal/service"
	"github.com/emptyhopes/employees_proxy/internal/validation"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type api struct {
	employeeValidation validation.InterfaceEmployeeValidation
	employeeConverter  converter.InterfaceEmployeeConverter
	employeeService    service.InterfaceEmployeeService
}

var _ definition.InterfaceEmployeeApi = (*api)(nil)

func NewEmployeeApi(
	employeeValidation validation.InterfaceEmployeeValidation,
	employeeConverter converter.InterfaceEmployeeConverter,
	employeeService service.InterfaceEmployeeService,
) *api {
	return &api{
		employeeValidation: employeeValidation,
		employeeConverter:  employeeConverter,
		employeeService:    employeeService,
	}
}

func (a *api) GetEmployeeById(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	employeeId := variables["id"]

	err := a.employeeValidation.GetEmployeeByIdValidation(employeeId)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(getErrorInJson(err.Error()))

		return
	}

	getEmployeeByIdDto := dto.NewGetEmployeeByIdDto(employeeId)

	employeeDto, err := a.employeeService.GetEmployeeById(getEmployeeByIdDto)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			response.Write(getErrorInJson(fmt.Sprintf("соотрудник не найден с employee_id: %s", employeeId)))

			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(getErrorInJson(err.Error()))

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(employeeDto)
}

func (a *api) CreateEmployee(response http.ResponseWriter, request *http.Request) {
	var createEmployeeDto dto.CreateEmployeeDto

	if err := json.NewDecoder(request.Body).Decode(&createEmployeeDto); err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(getErrorInJson(fmt.Sprintf("failed to parse request body: %s", err.Error())))

		return
	}

	defer request.Body.Close()

	err := a.employeeValidation.CreateEmployeeValidation(&createEmployeeDto)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(getErrorInJson(err.Error()))

		return
	}

	resultDto, err := a.employeeService.CreateEmployee(&createEmployeeDto)

	if err != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(getErrorInJson(err.Error()))

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	json.NewEncoder(response).Encode(resultDto)
}

func (a *api) NotFound(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNotFound)
	response.Write(getErrorInJson("not found"))
}

func (a *api) MethodNotAllowed(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusMethodNotAllowed)
	response.Write(getErrorInJson("method not allowed"))
}

func getErrorInJson(message string) []byte {
	type ErrorStruct struct {
		Error string `json:"error"`
	}

	errorStruct := &ErrorStruct{
		Error: message,
	}

	errJson, err := json.Marshal(errorStruct)

	if err != nil {
		return []byte(err.Error())
	}

	return errJson
}
