package repository

import dto "github.com/emptyhopes/employees_cron/internal/dto/employees"

type InterfaceEmployeeRepository interface {
	GetEmployeesWithoutConfirmation() (*dto.EmployeesDto, error)
	UpdateEmployeeConfirmation(*dto.EmployeeDto) error
}
