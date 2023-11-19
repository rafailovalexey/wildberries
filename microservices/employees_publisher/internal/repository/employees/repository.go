package employees

import (
	"github.com/emptyhopes/employees_publisher/internal/converter"
	dto "github.com/emptyhopes/employees_publisher/internal/dto/employees"
	model "github.com/emptyhopes/employees_publisher/internal/model/employees"
	definition "github.com/emptyhopes/employees_publisher/internal/repository"
	"github.com/google/uuid"
)

type repository struct {
	converter converter.InterfaceEmployeeConverter
}

var _ definition.InterfaceEmployeeRepository = (*repository)(nil)

func NewEmployeeRepository(converter converter.InterfaceEmployeeConverter) *repository {
	return &repository{
		converter: converter,
	}
}

func (r *repository) GetEmployee() (*dto.EmployeeDto, error) {
	employeeModel := model.NewEmployeeModel(
		uuid.New().String(),
	)

	employeeDto := r.converter.MapEmployeeModelToEmployeeDto(employeeModel)

	return employeeDto, nil
}
