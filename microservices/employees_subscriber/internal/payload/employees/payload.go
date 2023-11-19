package employees

type EmployeePayload struct {
	EmployeeId string
}

func NewEmployeePayload(
	employeeId string,
) *EmployeePayload {
	return &EmployeePayload{
		EmployeeId: employeeId,
	}
}
