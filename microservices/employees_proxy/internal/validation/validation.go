package validation

import dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"

type InterfaceEmployeeValidation interface {
	GetEmployeeByIdValidation(string) error
	CreateEmployeeValidation(*dto.CreateEmployeeDto) error
}
