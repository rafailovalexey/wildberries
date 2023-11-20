package service

import dto "github.com/emptyhopes/employees/internal/dto/employees"

type InterfaceEmployeeService interface {
	GetEmployeeById(*dto.GetEmployeeByIdDto) (*dto.EmployeeDto, error)
	CreateEmployee(employeeDto *dto.CreateEmployeeDto) error
}
