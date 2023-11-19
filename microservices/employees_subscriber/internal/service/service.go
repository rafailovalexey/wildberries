package service

import dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"

type InterfaceEmployeeService interface {
	CreateEmployee(*dto.EmployeeDto) error
}
