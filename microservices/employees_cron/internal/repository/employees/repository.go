package employees

import (
	"context"
	"github.com/emptyhopes/employees_cron/internal/converter"
	dto "github.com/emptyhopes/employees_cron/internal/dto/employees"
	model "github.com/emptyhopes/employees_cron/internal/model/employees"
	definition "github.com/emptyhopes/employees_cron/internal/repository"
	"github.com/emptyhopes/employees_cron/storage"
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

func (r *repository) GetEmployeesWithoutConfirmation() (*dto.EmployeesDto, error) {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	employeesWithoutConfirmationModel, err := getEmployeesWithoutConfirmation(pool)

	if err != nil {
		return nil, err
	}

	employeesWithoutConfirmationDto := r.employeeConverter.MapEmployeesModelToEmployeesDto(employeesWithoutConfirmationModel)

	return employeesWithoutConfirmationDto, nil
}

func (r *repository) UpdateEmployeeConfirmation(employeeDto *dto.EmployeeDto) error {
	r.rwmutex.Lock()
	defer r.rwmutex.Unlock()

	pool := r.database.GetPool()
	defer pool.Close()

	employeeModel := r.employeeConverter.MapEmployeeDtoToEmployeeModel(employeeDto)

	err := updateEmployeeConfirmation(pool, employeeModel)

	if err != nil {
		return err
	}

	return nil
}

func getEmployeesWithoutConfirmation(pool *pgxpool.Pool) (*model.EmployeesModel, error) {
	limit := 10
	confirmation := false

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
        FROM employees WHERE confirmation = $1
        LIMIT $2
    `

	rows, err := pool.Query(
		context.Background(),
		query,
		confirmation,
		limit,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make(model.EmployeesModel, 0, 10)

	for rows.Next() {
		item := model.EmployeeModel{}

		err = rows.Scan(
			&item.EmployeeId,
			&item.Confirmation,
			&item.Firstname,
			&item.Lastname,
			&item.Email,
			&item.PhoneNumber,
			&item.Address,
			&item.Position,
			&item.Department,
			&item.DateOfBirth,
			&item.HireDate,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return &items, nil
}

func updateEmployeeConfirmation(
	pool *pgxpool.Pool,
	employeeModel *model.EmployeeModel,
) error {
	query := `
        UPDATE employees 
		SET confirmation = $1
        WHERE employee_id = $2;
    `

	_, err := pool.Exec(
		context.Background(),
		query,
		true,
		employeeModel.EmployeeId,
	)

	if err != nil {
		return err
	}

	return nil
}
