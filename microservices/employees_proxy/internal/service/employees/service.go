package employees

import (
	"github.com/emptyhopes/employees_proxy/internal/client"
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	definition "github.com/emptyhopes/employees_proxy/internal/service"
)

type service struct {
	employeeClient client.InterfaceClientEmployees
}

var _ definition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(employeeClient client.InterfaceClientEmployees) *service {
	return &service{
		employeeClient: employeeClient,
	}
}

func (s *service) GetEmployeeById(getEmployeeByIdDto *dto.GetEmployeeByIdDto) (*dto.EmployeeDto, error) {
	employeeDto, err := s.employeeClient.GetEmployeeById(getEmployeeByIdDto)

	if err != nil {
		return nil, err
	}

	return employeeDto, err
}

func (s *service) CreateEmployee(createEmployeeDto *dto.CreateEmployeeDto) (*dto.ResultDto, error) {
	resultDto, err := s.employeeClient.CreateEmployee(createEmployeeDto)

	if err != nil {
		return nil, err
	}

	return resultDto, err
}
