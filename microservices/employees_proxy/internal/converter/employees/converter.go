package employees

import (
	definition "github.com/emptyhopes/employees_proxy/internal/converter"
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	request "github.com/emptyhopes/employees_proxy/internal/request/employees"
	"github.com/emptyhopes/employees_proxy/pkg/employees_v1"
	"time"
)

type converter struct{}

var _ definition.InterfaceEmployeeConverter = (*converter)(nil)

func NewEmployeeConverter() *converter {
	return &converter{}
}

func (c *converter) MapGetEmployeeByIdRequestToGetEmployeeByIdDto(request *request.GetEmployeeByIdRequest) *dto.GetEmployeeByIdDto {
	return dto.NewGetEmployeeByIdDto(
		request.EmployeeId,
	)
}

func (c *converter) MapCreateEmployeeRequestToCreateEmployeeDto(request *request.CreateEmployeeRequest) *dto.CreateEmployeeDto {
	return dto.NewCreateEmployeeDto(
		request.Firstname,
		request.Lastname,
		request.Email,
		request.PhoneNumber,
		request.Address,
		request.Position,
		request.Department,
		request.DateOfBirth,
		request.HireDate,
	)
}

func (c *converter) MapGetEmployeeByIdResponseToEmployeeDto(response *employees_v1.GetEmployeeByIdResponse) *dto.EmployeeDto {
	return dto.NewEmployeeDto(
		response.Employee.GetEmployeeId(),
		response.Employee.GetConfirmation(),
		response.Employee.GetFirstname(),
		response.Employee.GetLastname(),
		response.Employee.GetEmail(),
		response.Employee.GetPhoneNumber(),
		response.Employee.GetAddress(),
		response.Employee.GetPosition(),
		response.Employee.GetDepartment(),
		time.Unix(response.Employee.GetDateOfBirth().GetSeconds(), 0),
		time.Unix(response.Employee.GetHireDate().GetSeconds(), 0),
		time.Unix(response.Employee.GetCreatedAt().GetSeconds(), 0),
		time.Unix(response.Employee.GetUpdatedAt().GetSeconds(), 0),
	)
}

func (c *converter) MapResultResponseToResultDto(response *employees_v1.ResultResponse) *dto.ResultDto {
	return dto.NewResultDto(
		response.GetResult(),
	)
}
