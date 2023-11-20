package employees

type EmployeeModel struct {
	EmployeeId   string
	Test         string
	Confirmation bool
}

type EmployeesModel = []EmployeeModel

func NewEmployeeModel(
	employeeId string,
	test string,
	confirmation bool,
) *EmployeeModel {
	return &EmployeeModel{
		EmployeeId:   employeeId,
		Test:         test,
		Confirmation: confirmation,
	}
}
