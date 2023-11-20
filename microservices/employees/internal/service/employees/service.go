package employees

import (
	"encoding/json"
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	"github.com/emptyhopes/employees/internal/repository"
	defenition "github.com/emptyhopes/employees/internal/service"
	"github.com/emptyhopes/employees/storage"
)

type service struct {
	employeeRepository repository.InterfaceEmployeeRepository
	natsPublisher      storage.InterfaceNatsPublisher
}

var _ defenition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(
	employeeRepository repository.InterfaceEmployeeRepository,
	natsPublisher storage.InterfaceNatsPublisher,
) *service {
	return &service{
		employeeRepository: employeeRepository,
		natsPublisher:      natsPublisher,
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

func (s *service) CreateEmployee(createEmployeeDto *dto.CreateEmployeeDto) error {
	sc := s.natsPublisher.GetConnect()
	defer sc.Close()

	message, err := json.Marshal(createEmployeeDto)

	if err != nil {
		return err
	}

	subject := "create-employee"

	err = sc.Publish(subject, message)

	if err != nil {
		return err
	}

	return nil
}
