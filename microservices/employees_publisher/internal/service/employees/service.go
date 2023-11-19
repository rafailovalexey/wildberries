package employees

import (
	dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"
	"github.com/emptyhopes/employees_publisher/internal/repository"
	definition "github.com/emptyhopes/employees_publisher/internal/service"
)

type service struct {
	repository repository.InterfaceEmployeeRepository
}

var _ definition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(repository repository.InterfaceEmployeeRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetEmployee() (*dto.EmployeeDto, error) {
	employeeDto, err := s.repository.GetEmployee()

	if err != nil {
		return nil, err
	}

	return employeeDto, nil
}
