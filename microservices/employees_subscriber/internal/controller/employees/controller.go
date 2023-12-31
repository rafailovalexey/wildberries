package employees

import (
	"encoding/json"
	definition "github.com/emptyhopes/employees_subscriber/internal/controller"
	"github.com/emptyhopes/employees_subscriber/internal/converter"
	payload "github.com/emptyhopes/employees_subscriber/internal/payload/employees"
	"github.com/emptyhopes/employees_subscriber/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type controller struct {
	employeeService   service.InterfaceEmployeeService
	employeeConverter converter.InterfaceEmployeeConverter
}

var _ definition.InterfaceEmployeeController = (*controller)(nil)

func NewEmployeeController(
	employeeService service.InterfaceEmployeeService,
	employeeConverter converter.InterfaceEmployeeConverter,
) *controller {
	return &controller{
		employeeService:   employeeService,
		employeeConverter: employeeConverter,
	}
}

func (c *controller) CreateEmployee(
	message *stan.Msg,
) {
	var employeePayload payload.EmployeePayload

	err := json.Unmarshal(message.Data, &employeePayload)

	if err != nil {
		log.Printf("a parsing error occurred %v\n", err)

		return
	}

	employeeDto := c.employeeConverter.MapEmployeePayloadToEmployeeDto(&employeePayload)

	err = c.employeeService.CreateEmployee(employeeDto)

	if err != nil {
		log.Printf("an error occurred while creating an employee %v\n", err)

		return
	}

	log.Printf("employee successfully created with employee_id %s\n", employeeDto.EmployeeId)
}
