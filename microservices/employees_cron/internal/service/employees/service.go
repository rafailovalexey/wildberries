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
	log.Printf("cron started working\n")

	employeesWithoutConfirmationDto, err := s.employeeRepository.GetEmployeesWithoutConfirmation()

	if err != nil {
		log.Printf("an error occurred when retrieving a list of employees without a verified account %v\n", err)

		return
	}

	for _, dto := range *employeesWithoutConfirmationDto {
		err = s.employeeRepository.UpdateEmployeeConfirmation(&dto)

		if err != nil {
			log.Printf("an error occurred on an employee with the identifier employee_id %s %v\n", dto.EmployeeId, err)

			return
		}

		log.Printf("updated employee confirmation with employee_id %s", dto.EmployeeId)
	}

	log.Printf("cron has finished its job\n")
}
