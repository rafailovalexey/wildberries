package employees

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/employees_proxy/internal/api"
	"github.com/emptyhopes/employees_proxy/internal/converter"
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	"github.com/emptyhopes/employees_proxy/internal/service"
	"github.com/emptyhopes/employees_proxy/internal/validation"
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
	// Очень плохо выглядит нужно либо выносить в пакет отдельный, либо использовать уже какой-то отдельный пакет
	segments := strings.Split(request.URL.Path, "/")

	if len(segments) != 4 || segments[1] != "v1" || segments[2] != "employees" {
		http.Error(response, getErrorJson("неверный URL"), http.StatusBadRequest)

		return
	}

	id := segments[3]

	err := a.employeeValidation.GetEmployeeByIdValidation(id)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusBadRequest)

		return
	}

	getEmployeeByIdDto := dto.NewGetEmployeeByIdDto(id)

	employeeDto, err := a.employeeService.GetEmployeeById(getEmployeeByIdDto)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			http.Error(response, getErrorJson(fmt.Sprintf("соотрудник не найден с employee_id: %s", id)), http.StatusBadRequest)

			return
		}

		http.Error(response, getErrorJson(err.Error()), http.StatusBadRequest)

		return
	}

	employeeJson, err := json.Marshal(employeeDto)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusInternalServerError)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	_, err = response.Write(employeeJson)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusInternalServerError)

		return
	}
}

func (a *api) CreateEmployee(response http.ResponseWriter, request *http.Request) {
	var createEmployeeDto dto.CreateEmployeeDto

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&createEmployeeDto); err != nil {
		http.Error(response, getErrorJson(fmt.Sprintf("failed to parse request body: %s", err.Error())), http.StatusBadRequest)

		return
	}

	defer request.Body.Close()

	err := a.employeeValidation.CreateEmployeeValidation(&createEmployeeDto)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusBadRequest)

		return
	}

	resultDto, err := a.employeeService.CreateEmployee(&createEmployeeDto)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusBadRequest)

		return
	}

	resultJson, err := json.Marshal(resultDto)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusInternalServerError)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	_, err = response.Write(resultJson)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusInternalServerError)

		return
	}
}

/*
EmployeesHandler
Использовал парсинг URL, для того, чтобы добиться REST поведения
GetAllEmployees - /v1/employees
GetEmployeeById - /v1/employees/:id
*/
func (a *api) EmployeesHandler(response http.ResponseWriter, request *http.Request) {
	// Очень плохо выглядит нужно либо выносить в пакет отдельный, либо использовать уже какой-то отдельный пакет
	switch request.Method {
	case http.MethodGet:
		if strings.HasPrefix(request.URL.Path, "/v1/employees/") {
			a.GetEmployeeById(response, request)

			return
		}

		http.Error(response, getErrorJson("несуществующий url"), http.StatusNotFound)
	case http.MethodPost:
		a.CreateEmployee(response, request)
	default:
		http.Error(response, getErrorJson("несуществующий http метод"), http.StatusMethodNotAllowed)
	}
}

func getErrorJson(message string) string {
	return fmt.Sprintf("{\"error\":\"%s\"}", message)
}
