package provider

import (
	"github.com/emptyhopes/employees_publisher/internal/controller"
	"github.com/emptyhopes/employees_publisher/internal/converter"
	"github.com/emptyhopes/employees_publisher/internal/repository"
	"github.com/emptyhopes/employees_publisher/internal/service"
)

type InterfaceEmployeeProvider interface {
	GetEmployeeController() controller.InterfaceEmployeeController
	GetEmployeeService() service.InterfaceEmployeeService
	GetEmployeeRepository() repository.InterfaceEmployeeRepository
	GetEmployeeConverter() converter.InterfaceEmployeeConverter
}
