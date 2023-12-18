package grpc_server

import (
	"fmt"
	interceptor "github.com/emptyhopes/employees/cmd/grpc_server/intereptor"
	"github.com/emptyhopes/employees/cmd/grpc_server/middleware"
	"github.com/emptyhopes/employees/pkg/employees_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func Run(api employees_v1.EmployeesV1Server) {
	hostname := os.Getenv("HOSTNAME")

	port := os.Getenv("PORT")

	if port == "" {
		log.Panicf("specify the port")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Printf("%s\n", fmt.Sprintf("grpc server starts at address %s", address))

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Panicf("grpc server startup error %v", err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.TracecodeInterceptor(),
			interceptor.LoggingInterceptor(),
			middleware.AuthenticationTokenMiddleware(),
		),
	)

	reflection.Register(server)

	employees_v1.RegisterEmployeesV1Server(server, api)

	log.Printf("%s\n", fmt.Sprintf("grpc server is running at %s", address))

	err = server.Serve(listener)

	if err != nil {
		log.Panicf("grpc server startup error %v", err)
	}
}
