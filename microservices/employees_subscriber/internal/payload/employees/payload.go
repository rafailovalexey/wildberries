package employees

type EmployeePayload struct {
	Test string
}

func NewEmployeePayload(
	test string,
) *EmployeePayload {
	return &EmployeePayload{
		Test: test,
	}
}
