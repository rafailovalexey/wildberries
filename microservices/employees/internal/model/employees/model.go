package employees

type EmployeeModel struct {
	EmployeeId string
}

func NewEmployeeModel(
	employeeId string,
) *EmployeeModel {
	return &EmployeeModel{
		EmployeeId: employeeId,
	}
}
