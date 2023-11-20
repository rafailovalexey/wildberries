package employees

import (
	definition "github.com/emptyhopes/employees/internal/converter"
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	model "github.com/emptyhopes/employees/internal/model/employees"
	"github.com/emptyhopes/employees/pkg/employees_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		request.GetFirstname(),
		request.GetLastname(),
		request.GetEmail(),
		request.GetPhoneNumber(),
		request.GetAddress(),
		request.GetPosition(),
		request.GetDepartment(),
		request.GetDateOfBirth(),
		request.GetHireDate(),
	)
}

func (c *converter) MapEmployeeDtoToEmployeeResponse(
	dto *dto.EmployeeDto,
) *employees_v1.GetEmployeeByIdResponse {
	return &employees_v1.GetEmployeeByIdResponse{
		Employee: &employees_v1.Employee{
			EmployeeId:   dto.EmployeeId,
			Confirmation: false,
			Firstname:    dto.Firstname,
			Lastname:     dto.Lastname,
			Email:        dto.Email,
			PhoneNumber:  dto.PhoneNumber,
			Address:      dto.Address,
			Position:     dto.Position,
			Department:   dto.Department,
			DateOfBirth:  timestamppb.New(dto.DateOfBirth),
			HireDate:     timestamppb.New(dto.HireDate),
			CreatedAt:    timestamppb.New(dto.CreatedAt),
			UpdatedAt:    timestamppb.New(dto.UpdatedAt),
		},
	}
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
