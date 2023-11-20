package converter

import (
	dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"
	model "github.com/emptyhopes/employees_publisher/internal/model/employees"
)

type InterfaceEmployeeConverter interface {
	MapEmployeeModelToEmployeeDto(*model.EmployeeModel) *dto.EmployeeDto
}
