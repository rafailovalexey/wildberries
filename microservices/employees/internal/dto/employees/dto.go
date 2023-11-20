package employees

import (
	"time"
)

type GetEmployeeByIdDto struct {
	EmployeeId string
}

type CreateEmployeeDto struct {
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber string
	Address     string
	Position    string
	Department  string
	DateOfBirth time.Time
	HireDate    time.Time
}

type EmployeeDto struct {
	EmployeeId   string
	Confirmation bool
	Firstname    string
	Lastname     string
	Email        string
	PhoneNumber  string
	Address      string
	Position     string
	Department   string
	DateOfBirth  time.Time
	HireDate     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EmployeesDto = []EmployeeDto

func NewGetEmployeeByIdDto(
	employeeId string,
) *GetEmployeeByIdDto {
	return &GetEmployeeByIdDto{
		EmployeeId: employeeId,
	}
}

func NewCreateEmployeeDto(
	firstname string,
	lastname string,
	email string,
	phoneNumber string,
	address string,
	position string,
	department string,
	dateOfBirth time.Time,
	hireDate time.Time,
) *CreateEmployeeDto {
	return &CreateEmployeeDto{
		Firstname:   firstname,
		Lastname:    lastname,
		Email:       email,
		PhoneNumber: phoneNumber,
		Address:     address,
		Position:    position,
		Department:  department,
		DateOfBirth: dateOfBirth,
		HireDate:    hireDate,
	}
}

func NewEmployeeDto(
	employeeId string,
	confirmation bool,
	firstname string,
	lastname string,
	email string,
	phoneNumber string,
	address string,
	position string,
	department string,
	dateOfBirth time.Time,
	hireDate time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *EmployeeDto {
	return &EmployeeDto{
		EmployeeId:   employeeId,
		Confirmation: confirmation,
		Firstname:    firstname,
		Lastname:     lastname,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Address:      address,
		Position:     position,
		Department:   department,
		DateOfBirth:  dateOfBirth,
		HireDate:     hireDate,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
