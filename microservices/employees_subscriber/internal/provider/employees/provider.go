package employees

import (
	"github.com/emptyhopes/employees_subscriber/internal/controller"
	employeeController "github.com/emptyhopes/employees_subscriber/internal/controller/employees"
	"github.com/emptyhopes/employees_subscriber/internal/converter"
	employeeConverter "github.com/emptyhopes/employees_subscriber/internal/converter/employees"
	definition "github.com/emptyhopes/employees_subscriber/internal/provider"
	"github.com/emptyhopes/employees_subscriber/internal/repository"
	employeeRepository "github.com/emptyhopes/employees_subscriber/internal/repository/employees"
	"github.com/emptyhopes/employees_subscriber/internal/service"
	employeeService "github.com/emptyhopes/employees_subscriber/internal/service/employees"
	"github.com/emptyhopes/employees_subscriber/storage"
)

type provider struct {
	employeeController controller.InterfaceEmployeeController
	employeeService    service.InterfaceEmployeeService
	employeeRepository repository.InterfaceEmployeeRepository
	employeeConverter  converter.InterfaceEmployeeConverter
}

var _ definition.InterfaceEmployeeProvider = (*provider)(nil)

func NewEmployeeProvider() *provider {
	return &provider{}
}

func (p *provider) GetEmployeeController() controller.InterfaceEmployeeController {
	if p.employeeController == nil {
		p.employeeController = employeeController.NewEmployeeController(
			p.GetEmployeeService(),
		)
	}

	return p.employeeController
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
