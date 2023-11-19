package employees

import (
	"context"
	"github.com/emptyhopes/employees/internal/converter"
	"github.com/emptyhopes/employees/internal/service"
	"github.com/emptyhopes/employees/pkg/employees_v1"
)

type ImplementationEmployee struct {
	employees_v1.UnimplementedEmployeesV1Server
	service   service.InterfaceEmployeeService
	converter converter.InterfaceEmployeeConverter
}

func NewEmployeeImplementation(service service.InterfaceEmployeeService, converter converter.InterfaceEmployeeConverter) *ImplementationEmployee {
	return &ImplementationEmployee{
		service:   service,
		converter: converter,
	}
}

func (i *ImplementationEmployee) GetEmployeeById(_ context.Context, request *employees_v1.GetEmployeeByIdRequest) (*employees_v1.GetEmployeeByIdResponse, error) {
	getEmployeeByIdDto := i.converter.MapGetEmployeeByIdRequestToGetEmployeeByIdDto(request)

	employeeDto, err := i.service.GetEmployeeById(getEmployeeByIdDto)

	if err != nil {
		return nil, err
	}

	getEmployeeByIdResponse := i.converter.MapEmployeeDtoToEmployeeResponse(employeeDto)

	return getEmployeeByIdResponse, nil
}

func (i *ImplementationEmployee) mustEmbedUnimplementedEmployeesV1Server() {
	return
}
