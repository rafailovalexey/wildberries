package employees

import (
	definition "github.com/emptyhopes/employees_publisher/internal/converter"
	dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"
	model "github.com/emptyhopes/employees_publisher/internal/model/employees"
)

type converter struct{}

var _ definition.InterfaceEmployeeConverter = (*converter)(nil)

func NewEmployeeConverter() *converter {
	return &converter{}
}

func (c *converter) MapEmployeeModelToEmployeeDto(model *model.EmployeeModel) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		model.EmployeeId,
	)
}
