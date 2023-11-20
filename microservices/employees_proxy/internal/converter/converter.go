package converter

import (
	dto "github.com/emptyhopes/employees_cron/internal/dto/employees"
	model "github.com/emptyhopes/employees_cron/internal/model/employees"
)

type InterfaceEmployeeConverter interface {
	MapEmployeeDtoToEmployeeModel(*dto.EmployeeDto) *model.EmployeeModel
	MapEmployeesDtoToEmployeesModel(*dto.EmployeesDto) *model.EmployeesModel

	MapEmployeeModelToEmployeeDto(*model.EmployeeModel) *dto.EmployeeDto
	MapEmployeesModelToEmployeesDto(model *model.EmployeesModel) *dto.EmployeesDto
}
