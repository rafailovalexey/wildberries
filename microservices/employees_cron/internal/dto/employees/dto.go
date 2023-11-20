package employees

type EmployeeDto struct {
	EmployeeId   string
	Test         string
	Confirmation bool
}

type EmployeesDto = []EmployeeDto

func NewEmployeeDto(
	employeeId string,
	test string,
	confirmation bool,
) *EmployeeDto {
	return &EmployeeDto{
		EmployeeId:   employeeId,
		Test:         test,
		Confirmation: confirmation,
	}
}
