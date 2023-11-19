package employees

import (
	definition "github.com/emptyhopes/employees/internal/converter"
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	model "github.com/emptyhopes/employees/internal/model/employees"
	"github.com/emptyhopes/employees/pkg/employees_v1"
)

type converter struct{}

var _ definition.InterfaceEmployeeConverter = (*converter)(nil)

func NewEmployeeConverter() *converter {
	return &converter{}
}

func (c *converter) MapGetEmployeeByIdRequestToGetEmployeeByIdDto(response *employees_v1.GetEmployeeByIdRequest) *dto.GetEmployeeByIdDto {
	return dto.NewGetEmployeeByIdDto(
		response.GetEmployeeId(),
	)
}

func (c *converter) MapEmployeeDtoToEmployeeResponse(dto *dto.EmployeeDto) *employees_v1.GetEmployeeByIdResponse {
	return &employees_v1.GetEmployeeByIdResponse{
		Employee: &employees_v1.Employee{
			EmployeeId: dto.EmployeeId,
		},
	}
}

func (c *converter) MapEmployeeDtoToEmployeeModel(dto *dto.EmployeeDto) *model.EmployeeModel {
	return model.NewEmployeeModel(
		dto.EmployeeId,
	)
}

func (c *converter) MapEmployeeModelToEmployeeDto(model *model.EmployeeModel) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		model.EmployeeId,
	)
}
