package employees

type EmployeeDto struct {
	EmployeeId string
	Test       string
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
