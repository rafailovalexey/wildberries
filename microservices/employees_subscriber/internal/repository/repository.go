package repository

import dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"

type InterfaceEmployeeRepository interface {
	CreateEmployee(*dto.EmployeeDto) error
}
