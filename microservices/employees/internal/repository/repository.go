package repository

import dto "github.com/emptyhopes/employees/internal/dto/employees"

type InterfaceEmployeeRepository interface {
	GetEmployeeById(string) (*dto.EmployeeDto, error)
}
