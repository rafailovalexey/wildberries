package employees

import (
	"errors"
	"fmt"
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	definition "github.com/emptyhopes/employees_proxy/internal/validation"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

type validation struct{}

var _ definition.InterfaceEmployeeValidation = (*validation)(nil)

func NewEmployeeValidation() *validation {
	return &validation{}
}

func (v *validation) GetEmployeeByIdValidation(employeeId string) error {
	if err := isValidUuid(employeeId, "employee_id"); err != nil {
		return err
	}

	return nil
}

func (v *validation) CreateEmployeeValidation(createEmployeeDto *dto.CreateEmployeeDto) error {
	if err := isString(createEmployeeDto.Firstname, "firstname", 50); err != nil {
		return err
	}

	if err := isString(createEmployeeDto.Lastname, "lastname", 50); err != nil {
		return err
	}

	if err := isString(createEmployeeDto.Email, "email", 50); err != nil {
		return err
	}

	if err := isString(createEmployeeDto.PhoneNumber, "phone_number", 30); err != nil {
		return err
	}

	if err := isString(createEmployeeDto.Address, "address", 100); err != nil {
		return err
	}

	if err := isString(createEmployeeDto.Position, "position", 50); err != nil {
		return err
	}

	if err := isString(createEmployeeDto.Department, "department", 50); err != nil {
		return err
	}

	if err := isTime(createEmployeeDto.DateOfBirth, "date_of_birth"); err != nil {
		return err
	}

	if err := isTime(createEmployeeDto.HireDate, "hire_date"); err != nil {
		return err
	}

	return nil
}

func isValidUuid(id string, field string) error {
	result := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`).MatchString(id)

	if !result {
		return errors.New(fmt.Sprintf("%s is not uuid", field))
	}

	return nil
}

func isString(str string, field string, maxLength int) error {
	if reflect.TypeOf(str).Kind() != reflect.String {
		return errors.New(fmt.Sprintf("%s is not string", field))
	}

	if utf8.RuneCountInString(str) > maxLength {
		return errors.New(fmt.Sprintf("%s is string to big max length %d", field, maxLength))
	}

	return nil
}

func isTime(t time.Time, field string) error {
	if reflect.TypeOf(t) != reflect.TypeOf(time.Time{}) {
		return errors.New(fmt.Sprintf("%s is not a time", field))
	}

	return nil
}
