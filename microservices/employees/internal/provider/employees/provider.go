package employees

import (
	"github.com/emptyhopes/employees/internal/converter"
	employeeConverter "github.com/emptyhopes/employees/internal/converter/employees"
	"github.com/emptyhopes/employees/internal/implementation/employees"
	definition "github.com/emptyhopes/employees/internal/provider"
	"github.com/emptyhopes/employees/internal/repository"
	employeeRepository "github.com/emptyhopes/employees/internal/repository/employees"
	"github.com/emptyhopes/employees/internal/service"
	employeeService "github.com/emptyhopes/employees/internal/service/employees"
	"github.com/emptyhopes/employees/storage"
)

type provider struct {
	employeeImplementation *employees.ImplementationEmployee
	employeeService        service.InterfaceEmployeeService
	employeeRepository     repository.InterfaceEmployeeRepository
	employeeConverter      converter.InterfaceEmployeeConverter
}

var _ definition.InterfaceEmployeeProvider = (*provider)(nil)

func NewEmployeeProvider() *provider {
	return &provider{}
}

func (p *provider) GetEmployeeImplementation() *employees.ImplementationEmployee {
	if p.employeeImplementation == nil {
		p.employeeImplementation = employees.NewEmployeeImplementation(
			p.GetEmployeeService(),
			p.GetEmployeeConverter(),
		)
	}

	return p.employeeImplementation
}

func (p *provider) GetEmployeeService() service.InterfaceEmployeeService {
	if p.employeeService == nil {
		p.employeeService = employeeService.NewEmployeeService(
			p.GetEmployeeRepository(),
		)
	}

	return p.employeeService
}

func (p *provider) GetEmployeeRepository() repository.InterfaceEmployeeRepository {
	if p.employeeRepository == nil {
		p.employeeRepository = employeeRepository.NewEmployeeRepository(
			p.GetEmployeeConverter(),
			storage.NewDatabase(),
		)
	}

	return p.employeeRepository
}

func (p *provider) GetEmployeeConverter() converter.InterfaceEmployeeConverter {
	if p.employeeConverter == nil {
		p.employeeConverter = employeeConverter.NewEmployeeConverter()
	}

	return p.employeeConverter
}
