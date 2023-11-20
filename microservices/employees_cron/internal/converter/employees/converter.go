package employees

import (
	definition "github.com/emptyhopes/employees_cron/internal/converter"
	dto "github.com/emptyhopes/employees_cron/internal/dto/employees"
	model "github.com/emptyhopes/employees_cron/internal/model/employees"
)

type converter struct{}

var _ definition.InterfaceEmployeeConverter = (*converter)(nil)

func NewEmployeeConverter() *converter {
	return &converter{}
}

func (c *converter) MapEmployeeDtoToEmployeeModel(dto *dto.EmployeeDto) *model.EmployeeModel {
	return model.NewEmployeeModel(
		dto.EmployeeId,
		dto.Test,
		dto.Confirmation,
	)
}

func (c *converter) MapEmployeeModelToEmployeeDto(model *model.EmployeeModel) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		model.EmployeeId,
		model.Test,
		model.Confirmation,
	)
}

func (c *converter) MapEmployeesDtoToEmployeesModel(dtos *dto.EmployeesDto) *model.EmployeesModel {
	models := make(model.EmployeesModel, len(*dtos))

	for index, value := range *dtos {
		models[index] = *c.MapEmployeeDtoToEmployeeModel(&value)
	}

	return &models
}

func (c *converter) MapEmployeesModelToEmployeesDto(models *model.EmployeesModel) *dto.EmployeesDto {
	dtos := make(dto.EmployeesDto, len(*models))

	for index, value := range *models {
		dtos[index] = *c.MapEmployeeModelToEmployeeDto(&value)
	}

	return &dtos
}
