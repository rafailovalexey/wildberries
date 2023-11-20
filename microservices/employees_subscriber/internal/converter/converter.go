package converter

import (
	dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"
	model "github.com/emptyhopes/employees_subscriber/internal/model/employees"
	payload "github.com/emptyhopes/employees_subscriber/internal/payload/employees"
)

type InterfaceEmployeeConverter interface {
	MapEmployeePayloadToEmployeeDto(*payload.EmployeePayload) *dto.EmployeeDto

	MapEmployeeDtoToEmployeeModel(*dto.EmployeeDto) *model.EmployeeModel

	MapEmployeeModelToEmployeeDto(*model.EmployeeModel) *dto.EmployeeDto
}
