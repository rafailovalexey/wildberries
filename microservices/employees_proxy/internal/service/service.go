package service

import dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"

type InterfaceEmployeeService interface {
	GetEmployeeById(*dto.GetEmployeeByIdDto) (*dto.EmployeeDto, error)
	CreateEmployee(*dto.CreateEmployeeDto) (*dto.ResultDto, error)
}
