package employees

import (
	definition "github.com/emptyhopes/employees_subscriber/internal/converter"
	dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"
	model "github.com/emptyhopes/employees_subscriber/internal/model/employees"
	payload "github.com/emptyhopes/employees_subscriber/internal/payload/employees"
	"github.com/google/uuid"
)

type converter struct{}

var _ definition.InterfaceEmployeeConverter = (*converter)(nil)

func NewEmployeeConverter() *converter {
	return &converter{}
}

func (c *converter) MapEmployeePayloadToEmployeeDto(payload *payload.EmployeePayload) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		uuid.New().String(),
		payload.Confirmation,
		payload.Firstname,
		payload.Lastname,
		payload.Email,
		payload.PhoneNumber,
		payload.Address,
		payload.Position,
		payload.Department,
		payload.DateOfBirth,
		payload.HireDate,
		payload.CreatedAt,
		payload.UpdatedAt,
	)
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
