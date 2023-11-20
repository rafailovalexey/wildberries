package employees

type EmployeeModel struct {
	EmployeeId string
	Test       string
}

func NewEmployeeModel(
	employeeId string,
	test string,
) *EmployeeModel {
	return &EmployeeModel{
		EmployeeId: employeeId,
		Test:       test,
	}
}
