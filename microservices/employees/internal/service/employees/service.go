package employees

import (
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	"github.com/emptyhopes/employees/internal/repository"
	defenition "github.com/emptyhopes/employees/internal/service"
)

type service struct {
	repository repository.InterfaceEmployeeRepository
}

var _ defenition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(repository repository.InterfaceEmployeeRepository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetEmployeeById(getEmployeeByIdDto *dto.GetEmployeeByIdDto) (*dto.EmployeeDto, error) {
	employeeDto, err := s.repository.GetEmployeeById(getEmployeeByIdDto.EmployeeId)

	if err != nil {
		return nil, err
	}

	return employeeDto, nil
}
