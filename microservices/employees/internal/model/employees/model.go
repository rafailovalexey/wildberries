package employees

import (
	"time"
)

type EmployeeModel struct {
	EmployeeId  string
	Firstname   string
	Lastname    string
	Email       string
	PhoneNumber string
	Address     string
	Position    string
	Department  string
	DateOfBirth time.Time
	HireDate    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewEmployeeModel(
	employeeId string,
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
) *EmployeeModel {
	return &EmployeeModel{
		EmployeeId:  employeeId,
		Firstname:   firstname,
		Lastname:    lastname,
		Email:       email,
		PhoneNumber: phoneNumber,
		Address:     address,
		Position:    position,
		Department:  department,
		DateOfBirth: dateOfBirth,
		HireDate:    hireDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
