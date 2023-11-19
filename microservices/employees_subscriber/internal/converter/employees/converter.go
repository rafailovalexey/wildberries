package employees

import (
	definition "github.com/emptyhopes/employees_subscriber/internal/converter"
	dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"
	model "github.com/emptyhopes/employees_subscriber/internal/model/employees"
	payload "github.com/emptyhopes/employees_subscriber/internal/payload/employees"
)

type converter struct{}

var _ definition.InterfaceEmployeeConverter = (*converter)(nil)

func NewEmployeeConverter() *converter {
	return &converter{}
}

func (c *converter) MapEmployeePayloadToEmployeeDto(payload *payload.EmployeePayload) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		payload.EmployeeId,
	)
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
