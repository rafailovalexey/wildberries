package employees

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/employees_publisher/internal/controller"
	"github.com/emptyhopes/employees_publisher/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type controller struct {
	employeeService service.InterfaceEmployeeService
}

var _ definition.InterfaceEmployeeController = (*controller)(nil)

func NewEmployeeController(
	employeeService service.InterfaceEmployeeService,
) *controller {
	return &controller{
		employeeService: employeeService,
	}
}

func (c *controller) PublishEmployee(
	sc stan.Conn,
	subject string,
) {
	employeeDto, err := c.employeeService.GetEmployee()

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	message, err := json.Marshal(employeeDto)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	err = sc.Publish(subject, []byte(message))

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	if err == nil {
		fmt.Printf("опубликовал сообщение с employee_id: %s\n", employeeDto.EmployeeId)
	}
}
