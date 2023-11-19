package employees

import (
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	"github.com/emptyhopes/employees/internal/repository"
	defenition "github.com/emptyhopes/employees/internal/service"
)

type service struct {
	employeeRepository repository.InterfaceEmployeeRepository
}

var _ defenition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(
	employeeRepository repository.InterfaceEmployeeRepository,
) *service {
	return &service{
		employeeRepository: employeeRepository,
	}
}

func (s *service) GetEmployeeById(
	getEmployeeByIdDto *dto.GetEmployeeByIdDto,
) (*dto.EmployeeDto, error) {
	employeeDto, err := s.employeeRepository.GetEmployeeById(getEmployeeByIdDto.EmployeeId)

	if err != nil {
		return nil, err
	}

	return employeeDto, nil
}
