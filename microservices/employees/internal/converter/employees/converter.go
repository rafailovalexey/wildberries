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

func (c *converter) MapGetEmployeeByIdRequestToGetEmployeeByIdDto(
	response *employees_v1.GetEmployeeByIdRequest,
) *dto.GetEmployeeByIdDto {
	return dto.NewGetEmployeeByIdDto(
		response.GetEmployeeId(),
	)
}

func (c *converter) MapCreateEmployeeRequestToCreateEmployeeDto(request *employees_v1.CreateEmployeeRequest) *dto.CreateEmployeeDto {
	return dto.NewCreateEmployeeDto(
		request.GetTest(),
	)
}

func (c *converter) MapEmployeeDtoToEmployeeResponse(
	dto *dto.EmployeeDto,
) *employees_v1.GetEmployeeByIdResponse {
	return &employees_v1.GetEmployeeByIdResponse{
		Employee: &employees_v1.Employee{
			EmployeeId: dto.EmployeeId,
		},
	}
}

func (c *converter) MapEmployeeDtoToEmployeeModel(
	dto *dto.EmployeeDto,
) *model.EmployeeModel {
	return model.NewEmployeeModel(
		dto.EmployeeId,
		dto.Test,
	)
}

func (c *converter) MapEmployeeModelToEmployeeDto(
	model *model.EmployeeModel,
) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		model.EmployeeId,
		model.Test,
	)
}
