package employees

import (
	"context"
	"github.com/emptyhopes/employees/internal/converter"
	"github.com/emptyhopes/employees/internal/service"
	"github.com/emptyhopes/employees/pkg/employees_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImplementationEmployee struct {
	employees_v1.UnimplementedEmployeesV1Server
	employeeService   service.InterfaceEmployeeService
	employeeConverter converter.InterfaceEmployeeConverter
}

func NewEmployeeImplementation(
	employeeService service.InterfaceEmployeeService,
	employeeConverter converter.InterfaceEmployeeConverter,
) *ImplementationEmployee {
	return &ImplementationEmployee{
		employeeService:   employeeService,
		employeeConverter: employeeConverter,
	}
}

func (i *ImplementationEmployee) GetEmployeeById(
	_ context.Context,
	request *employees_v1.GetEmployeeByIdRequest,
) (*employees_v1.GetEmployeeByIdResponse, error) {
	getEmployeeByIdDto := i.employeeConverter.MapGetEmployeeByIdRequestToGetEmployeeByIdDto(request)

	employeeDto, err := i.employeeService.GetEmployeeById(getEmployeeByIdDto)

	if err != nil {
		return nil, err
	}

	getEmployeeByIdResponse := i.employeeConverter.MapEmployeeDtoToEmployeeResponse(employeeDto)

	return getEmployeeByIdResponse, nil
}

func (i *ImplementationEmployee) CreateEmployee(
	_ context.Context,
	request *employees_v1.CreateEmployeeRequest,
) (*employees_v1.ResultResponse, error) {
	createEmployeeDto := i.employeeConverter.MapCreateEmployeeRequestToCreateEmployeeDto(request)

	err := i.employeeService.CreateEmployee(createEmployeeDto)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &employees_v1.ResultResponse{Result: true}, nil
}

func (i *ImplementationEmployee) mustEmbedUnimplementedEmployeesV1Server() {
	return
}
