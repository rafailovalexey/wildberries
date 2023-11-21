package employees

import "time"

type GetEmployeeByIdDto struct {
	EmployeeId string `json:"employee_id"`
}

type CreateEmployeeDto struct {
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Position    string    `json:"position"`
	Department  string    `json:"department"`
	DateOfBirth time.Time `json:"date_of_birth"`
	HireDate    time.Time `json:"hire_date"`
}

type EmployeeDto struct {
	EmployeeId   string    `json:"employee_id"`
	Confirmation bool      `json:"confirmation"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
	Position     string    `json:"position"`
	Department   string    `json:"department"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	HireDate     time.Time `json:"hire_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EmployeesDto = []EmployeeDto

type ResultDto struct {
	Result bool `json:"result"`
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

func NewResultDto(
	result bool,
) *ResultDto {
	return &ResultDto{
		Result: result,
	}
}
