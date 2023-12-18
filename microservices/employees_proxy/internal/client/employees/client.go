package employees

import (
	"context"
	definition "github.com/emptyhopes/employees_proxy/internal/client"
	"github.com/emptyhopes/employees_proxy/internal/converter"
	dto "github.com/emptyhopes/employees_proxy/internal/dto/employees"
	"github.com/emptyhopes/employees_proxy/pkg/employees_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
)

type client struct {
	employeeConverter converter.InterfaceEmployeeConverter
}

var _ definition.InterfaceClientEmployees = (*client)(nil)

func NewEmployeeClient(
	employeeConverter converter.InterfaceEmployeeConverter,
) *client {
	return &client{
		employeeConverter: employeeConverter,
	}
}

func (c *client) GetEmployeeById(getEmployeeByIdDto *dto.GetEmployeeByIdDto) (*dto.EmployeeDto, error) {
	// Тут надо бы сделать сборщик клиента с учётом того, что клиентов может быть много
	url := os.Getenv("EMPLOYEES_URL")

	if url == "" {
		log.Panicf("specify employees url")
	}

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Panicf("error: %v\n", err)
	}

	cl := employees_v1.NewEmployeesV1Client(conn)

	authenticationTokenHeader := os.Getenv("EMPLOYEES_AUTHENTICATION_TOKEN_HEADER")

	if authenticationTokenHeader == "" {
		log.Panicf("specify employees header token authentication")
	}

	authenticationToken := os.Getenv("EMPLOYEES_AUTHENTICATION_TOKEN")

	if authenticationToken == "" {
		log.Panicf("specify employees token authentication")
	}

	md := metadata.Pairs(
		authenticationTokenHeader, authenticationToken,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	request := &employees_v1.GetEmployeeByIdRequest{
		EmployeeId: getEmployeeByIdDto.EmployeeId,
	}

	response, err := cl.GetEmployeeById(ctx, request)

	if err != nil {
		return nil, err
	}

	employeeDto := c.employeeConverter.MapGetEmployeeByIdResponseToEmployeeDto(response)

	return employeeDto, err
}

func (c *client) CreateEmployee(createEmployeeDto *dto.CreateEmployeeDto) (*dto.ResultDto, error) {
	// Тут надо бы сделать сборщик клиента с учётом того, что клиентов может быть много
	url := os.Getenv("EMPLOYEES_URL")

	if url == "" {
		log.Panicf("specify employees url")
	}

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Panicf("error: %v\n", err)
	}

	cl := employees_v1.NewEmployeesV1Client(conn)

	authenticationTokenHeader := os.Getenv("EMPLOYEES_AUTHENTICATION_TOKEN_HEADER")

	if authenticationTokenHeader == "" {
		log.Panicf("specify employees header token authentication")
	}

	authenticationToken := os.Getenv("EMPLOYEES_AUTHENTICATION_TOKEN")

	if authenticationToken == "" {
		log.Panicf("specify employees token authentication")
	}

	md := metadata.Pairs(
		authenticationTokenHeader, authenticationToken,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	request := &employees_v1.CreateEmployeeRequest{
		Firstname:   createEmployeeDto.Firstname,
		Lastname:    createEmployeeDto.Lastname,
		Email:       createEmployeeDto.Email,
		PhoneNumber: createEmployeeDto.PhoneNumber,
		Address:     createEmployeeDto.Address,
		Position:    createEmployeeDto.Position,
		Department:  createEmployeeDto.Department,
		DateOfBirth: timestamppb.New(createEmployeeDto.DateOfBirth),
		HireDate:    timestamppb.New(createEmployeeDto.HireDate),
	}

	response, err := cl.CreateEmployee(ctx, request)

	if err != nil {
		return nil, err
	}

	resultDto := c.employeeConverter.MapResultResponseToResultDto(response)

	return resultDto, err
}
