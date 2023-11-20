package employees

import (
	dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"
	"github.com/emptyhopes/employees_publisher/internal/repository"
	definition "github.com/emptyhopes/employees_publisher/internal/service"
)

type service struct {
	employeeRepository repository.InterfaceEmployeeRepository
}

var _ definition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(
	employeeRepository repository.InterfaceEmployeeRepository,
) *service {
	return &service{
		employeeRepository: employeeRepository,
	}
}

func (s *service) GetEmployee() (*dto.EmployeeDto, error) {
	employeeDto, err := s.employeeRepository.GetEmployee()

	if err != nil {
		return nil, err
	}

	return employeeDto, nil
}
