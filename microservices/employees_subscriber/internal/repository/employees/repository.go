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
        )
        VALUES (
        	$1,
            $2,
            $3,
            $4,
            $5,
            $6,
            $7,
            $8,
            $9,
            $10,
            $11,
            $12,
            $13
        );
    `

	_, err := pool.Exec(
		context.Background(),
		query,
		model.EmployeeId,
		model,
	)

	if err != nil {
		return err
	}

	return nil
}
