package employees

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/employees_subscriber/internal/controller"
	"github.com/emptyhopes/employees_subscriber/internal/converter"
	payload "github.com/emptyhopes/employees_subscriber/internal/payload/employees"
	"github.com/emptyhopes/employees_subscriber/internal/service"
	"github.com/nats-io/stan.go"
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
		fmt.Printf("произошла ошибка парсинга %v\n", err)

		return
	}

	employeeDto := c.employeeConverter.MapEmployeePayloadToEmployeeDto(&employeePayload)

	err = c.employeeService.CreateEmployee(employeeDto)

	if err != nil {
		fmt.Printf("произошла ошибка при создание сотрудника %v\n", err)

		return
	}

	fmt.Printf("сотрудник успешно создан с employee_id: %s", employeeDto.EmployeeId)
}
