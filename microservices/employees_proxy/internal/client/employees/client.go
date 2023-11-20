package employees

import (
	"context"
	"github.com/emptyhopes/employees_proxy/pkg/employees_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

func GetEmployeeById(employeeId string) *employees_v1.GetEmployeeByIdResponse {
	url := os.Getenv("EMPLOYEES_URL")

	if url == "" {
		log.Fatalf("укажите employees url")
	}

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	client := employees_v1.NewEmployeesV1Client(conn)

	authenticationTokenHeader := os.Getenv("EMPLOYEES_AUTHENTICATION_TOKEN_HEADER")

	if authenticationTokenHeader == "" {
		log.Fatalf("укажите employees header token authentication")
	}

	authenticationToken := os.Getenv("EMPLOYEES_AUTHENTICATION_TOKEN")

	if authenticationToken == "" {
		log.Fatalf("укажите employees token authentication")
	}

	md := metadata.Pairs(
		authenticationTokenHeader, authenticationToken,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	request := &employees_v1.GetEmployeeByIdRequest{
		EmployeeId: employeeId,
	}

	response, err := client.GetEmployeeById(ctx, request)

	return response
}
