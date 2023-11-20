package employees

type GetEmployeeByIdDto struct {
	EmployeeId string
}

type CreateEmployeeDto struct {
	Test string
}

type EmployeeDto struct {
	EmployeeId string
	Test       string
}

func NewGetEmployeeByIdDto(
	employeeId string,
) *GetEmployeeByIdDto {
	return &GetEmployeeByIdDto{
		EmployeeId: employeeId,
	}
}

func NewCreateEmployeeDto(
	test string,
) *CreateEmployeeDto {
	return &CreateEmployeeDto{
		Test: test,
	}
}

func NewEmployeeDto(
	employeeId string,
	test string,
) *EmployeeDto {
	return &EmployeeDto{
		EmployeeId: employeeId,
		Test:       test,
	}
}
