package service

import dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"

type InterfaceEmployeeService interface {
	GetEmployee() (*dto.EmployeeDto, error)
}
