package provider

import (
	"github.com/emptyhopes/employees_cron/internal/converter"
	"github.com/emptyhopes/employees_cron/internal/repository"
	"github.com/emptyhopes/employees_cron/internal/service"
)

type InterfaceEmployeeProvider interface {
	GetEmployeeService() service.InterfaceEmployeeService
	GetEmployeeRepository() repository.InterfaceEmployeeRepository
	GetEmployeeConverter() converter.InterfaceEmployeeConverter
}
