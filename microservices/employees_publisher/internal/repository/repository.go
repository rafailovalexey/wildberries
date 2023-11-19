package repository

import dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"

type InterfaceEmployeeRepository interface {
	GetEmployee() (*dto.EmployeeDto, error)
}
