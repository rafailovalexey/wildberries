package converter

import (
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	model "github.com/emptyhopes/employees/internal/model/employees"
	"github.com/emptyhopes/employees/pkg/employees_v1"
)

type InterfaceEmployeeConverter interface {
	MapGetEmployeeByIdRequestToGetEmployeeByIdDto(*employees_v1.GetEmployeeByIdRequest) *dto.GetEmployeeByIdDto

	MapEmployeeDtoToEmployeeModel(*dto.EmployeeDto) *model.EmployeeModel
	MapEmployeeDtoToEmployeeResponse(*dto.EmployeeDto) *employees_v1.GetEmployeeByIdResponse

	MapEmployeeModelToEmployeeDto(*model.EmployeeModel) *dto.EmployeeDto
}
