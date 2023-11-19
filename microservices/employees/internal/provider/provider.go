package provider

import (
	"github.com/emptyhopes/employees/internal/converter"
	"github.com/emptyhopes/employees/internal/implementation/employees"
	"github.com/emptyhopes/employees/internal/repository"
	"github.com/emptyhopes/employees/internal/service"
)

type InterfaceEmployeeProvider interface {
	GetEmployeeImplementation() *employees.ImplementationEmployee
	GetEmployeeService() service.InterfaceEmployeeService
	GetEmployeeRepository() repository.InterfaceEmployeeRepository
	GetEmployeeConverter() converter.InterfaceEmployeeConverter
}
