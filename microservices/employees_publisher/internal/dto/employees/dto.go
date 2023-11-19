package employees

type EmployeeDto struct {
	EmployeeId string
}

func NewEmployeeDto(
	employeeId string,
) *EmployeeDto {
	return &EmployeeDto{
		EmployeeId: employeeId,
	}
}
