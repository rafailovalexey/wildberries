package employees

import (
	"context"
	"github.com/emptyhopes/employees/internal/converter"
	dto "github.com/emptyhopes/employees/internal/dto/employees"
	model "github.com/emptyhopes/employees/internal/model/employees"
	definition "github.com/emptyhopes/employees/internal/repository"
	"github.com/emptyhopes/employees/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type repository struct {
	employeeConverter converter.InterfaceEmployeeConverter
	database          storage.InterfaceDatabase
	rwmutex           sync.RWMutex
}

var _ definition.InterfaceEmployeeRepository = (*repository)(nil)

func NewEmployeeRepository(
	employeeConverter converter.InterfaceEmployeeConverter,
	database storage.InterfaceDatabase,
) *repository {
	return &repository{
		employeeConverter: employeeConverter,
		database:          database,
		rwmutex:           sync.RWMutex{},
	}
}

func (r *repository) GetEmployeeById(
	employeeId string,
) (*dto.EmployeeDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	employeeModel, err := getEmployee(pool, employeeId)
	if err != nil {
		return nil, err
	}

	employeeDto := r.employeeConverter.MapEmployeeModelToEmployeeDto(employeeModel)

	return employeeDto, nil
}

func getEmployee(
	pool *pgxpool.Pool,
	employeeId string,
) (*model.EmployeeModel, error) {
	query := `
        SELECT
			employee_id,
			confirmation,
			firstname,
			lastname,
			email,
			phone_number,
			address,
			position,
			department,
			date_of_birth,
			hire_date,
			created_at,
			updated_at
        FROM employees
        WHERE employee_id = $1
    `

	employee := model.EmployeeModel{}

	err := pool.QueryRow(
		context.Background(),
		query,
		employeeId,
	).Scan(
		&employee.EmployeeId,
		&employee.Confirmation,
		&employee.Firstname,
		&employee.Lastname,
		&employee.Email,
		&employee.PhoneNumber,
		&employee.Address,
		&employee.Position,
		&employee.Department,
		&employee.DateOfBirth,
		&employee.HireDate,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}
