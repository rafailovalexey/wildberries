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
	converter converter.InterfaceEmployeeConverter
	database  storage.DatabaseInterface
	rwmutex   sync.RWMutex
}

var _ definition.InterfaceEmployeeRepository = (*repository)(nil)

func NewEmployeeRepository(converter converter.InterfaceEmployeeConverter, database storage.DatabaseInterface) *repository {
	return &repository{
		converter: converter,
		database:  database,
	}
}

func (r *repository) GetEmployeeById(employeeId string) (*dto.EmployeeDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	employeeModel, err := getEmployee(pool, employeeId)
	if err != nil {
		return nil, err
	}

	employeeDto := r.converter.MapEmployeeModelToEmployeeDto(employeeModel)

	return employeeDto, nil
}

func getEmployee(pool *pgxpool.Pool, employeeId string) (*model.EmployeeModel, error) {
	query := `
        SELECT
           	employee_id,         
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
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}
