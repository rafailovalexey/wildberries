package provider

import (
	"github.com/emptyhopes/employees_subscriber/internal/controller"
	"github.com/emptyhopes/employees_subscriber/internal/converter"
	"github.com/emptyhopes/employees_subscriber/internal/repository"
	"github.com/emptyhopes/employees_subscriber/internal/service"
)

type InterfaceEmployeeProvider interface {
	GetEmployeeController() controller.InterfaceEmployeeController
	GetEmployeeService() service.InterfaceEmployeeService
	GetEmployeeRepository() repository.InterfaceEmployeeRepository
	GetEmployeeConverter() converter.InterfaceEmployeeConverter
}
