package employees

import (
	"context"
	"github.com/emptyhopes/employees_subscriber/internal/converter"
	dto "github.com/emptyhopes/employees_subscriber/internal/dto/employees"
	model "github.com/emptyhopes/employees_subscriber/internal/model/employees"
	definition "github.com/emptyhopes/employees_subscriber/internal/repository"
	"github.com/emptyhopes/employees_subscriber/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
)

type repository struct {
	employeeConverter converter.InterfaceEmployeeConverter
	database          storage.DatabaseInterface
	rwmutex           sync.RWMutex
}

var _ definition.InterfaceEmployeeRepository = (*repository)(nil)

func NewEmployeeRepository(
	employeeConverter converter.InterfaceEmployeeConverter,
	database storage.DatabaseInterface,
) *repository {
	return &repository{
		employeeConverter: employeeConverter,
		database:          database,
		rwmutex:           sync.RWMutex{},
	}
}

func (r *repository) CreateEmployee(
	employeeDto *dto.EmployeeDto,
) error {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	employeeModel := r.employeeConverter.MapEmployeeDtoToEmployeeModel(employeeDto)

	err := insertEmployee(pool, employeeModel)

	if err != nil {
		return err
	}

	return nil
}

func insertEmployee(
	pool *pgxpool.Pool,
	model *model.EmployeeModel,
) error {
	query := `
        INSERT INTO employees (
        	employee_id
        )
        VALUES (
        	$1
        );
    `

	_, err := pool.Query(
		context.Background(),
		query,
		model.EmployeeId,
	)

	if err != nil {
		return err
	}

	return nil
}
