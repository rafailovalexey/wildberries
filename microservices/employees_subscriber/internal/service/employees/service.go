package employees

import (
	dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"
	"github.com/emptyhopes/employees_subscriber/internal/repository"
	definition "github.com/emptyhopes/employees_subscriber/internal/service"
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

func (s *service) CreateEmployee(
	employeeDto *dto.EmployeeDto,
) error {
	err := s.employeeRepository.CreateEmployee(employeeDto)

	if err != nil {
		return err
	}

	return nil
}
