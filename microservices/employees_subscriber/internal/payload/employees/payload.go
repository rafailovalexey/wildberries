package employees

import (
	"time"
)

type EmployeePayload struct {
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

func NewEmployeePayload(
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
) *EmployeePayload {
	return &EmployeePayload{
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
