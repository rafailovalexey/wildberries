package employees

import (
	"github.com/emptyhopes/employees_cron/internal/repository"
	definition "github.com/emptyhopes/employees_cron/internal/service"
	"log"
)

type service struct {
	employeeRepository repository.InterfaceEmployeeRepository
}

var _ definition.InterfaceEmployeeService = (*service)(nil)

func NewEmployeeService(
	employeeRepository repository.InterfaceEmployeeRepository,
) *service {
	return &service{
		employeeRepository: employeeRepository,
	}
}

func (s *service) UpdateEmployeeWithoutConfirmation() {
	log.Printf("%s\n", "cron начал свою работу")

	employeesWithoutConfirmationDto, err := s.employeeRepository.GetEmployeesWithoutConfirmation()

	if err != nil {
		log.Printf("произошла ошибка при получение списка соотрудников без подтверждённого аккаунта %v\n", err)

		return
	}

	for _, dto := range *employeesWithoutConfirmationDto {
		err = s.employeeRepository.UpdateEmployeeConfirmation(&dto)

		if err != nil {
			log.Printf("произошла ошибка на сотрдунике с идентификатором employee_id: %s %v\n", dto.EmployeeId, err)

			return
		}

		log.Printf("обновил подтверждение сотрудника с идентификатором employee_id: %s", dto.EmployeeId)
	}

	log.Printf("%s\n", "cron завершил свою работу")
}
