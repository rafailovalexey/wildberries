package employees

import "time"

type GetEmployeeByIdRequest struct {
	EmployeeId string
}

type CreateEmployeeRequest struct {
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

func NewGetEmployeeByIdRequest(
	employeeId string,
) *GetEmployeeByIdRequest {
	return &GetEmployeeByIdRequest{
		EmployeeId: employeeId,
	}
}

func NewCreateEmployeeRequest(
	firstname string,
	lastname string,
	email string,
	phoneNumber string,
	address string,
	position string,
	department string,
	dateOfBirth time.Time,
	hireDate time.Time,
) *CreateEmployeeRequest {
	return &CreateEmployeeRequest{
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
