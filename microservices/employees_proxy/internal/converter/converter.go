package converter

import (
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	request "github.com/emptyhopes/employees_proxy/internal/request/employees"
	"github.com/emptyhopes/employees_proxy/pkg/employees_v1"
)

type InterfaceEmployeeConverter interface {
	MapGetEmployeeByIdRequestToGetEmployeeByIdDto(*request.GetEmployeeByIdRequest) *dto.GetEmployeeByIdDto
	MapCreateEmployeeRequestToCreateEmployeeDto(*request.CreateEmployeeRequest) *dto.CreateEmployeeDto

	MapGetEmployeeByIdResponseToEmployeeDto(*employees_v1.GetEmployeeByIdResponse) *dto.EmployeeDto
	MapResultResponseToResultDto(*employees_v1.ResultResponse) *dto.ResultDto
}
