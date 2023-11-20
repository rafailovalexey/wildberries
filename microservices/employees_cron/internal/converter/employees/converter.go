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

func (c *converter) MapEmployeeDtoToEmployeeModel(
	dto *dto.EmployeeDto,
) *model.EmployeeModel {
	return model.NewEmployeeModel(
		dto.EmployeeId,
		dto.Confirmation,
		dto.Firstname,
		dto.Lastname,
		dto.Email,
		dto.PhoneNumber,
		dto.Address,
		dto.Position,
		dto.Department,
		dto.DateOfBirth,
		dto.HireDate,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}

func (c *converter) MapEmployeeModelToEmployeeDto(
	model *model.EmployeeModel,
) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		model.EmployeeId,
		model.Confirmation,
		model.Firstname,
		model.Lastname,
		model.Email,
		model.PhoneNumber,
		model.Address,
		model.Position,
		model.Department,
		model.DateOfBirth,
		model.HireDate,
		model.CreatedAt,
		model.UpdatedAt,
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
