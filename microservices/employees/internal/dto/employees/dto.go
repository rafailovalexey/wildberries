package employees

type GetEmployeeByIdDto struct {
	EmployeeId string
}

type EmployeeDto struct {
	EmployeeId string
}

func NewGetEmployeeByIdDto(
	employeeId string,
) *GetEmployeeByIdDto {
	return &GetEmployeeByIdDto{
		EmployeeId: employeeId,
	}
}

func NewEmployeeDto(
	employeeId string,
) *EmployeeDto {
	return &EmployeeDto{
		EmployeeId: employeeId,
	}
}
