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
		log.Fatalf("укажите порт")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Println(fmt.Sprintf("GRPC сервер запускается по адресу %s", address))

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("ошибка запуска grpc сервера %v", err)
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

	log.Println(fmt.Sprintf("GRPC сервер запущен по адресу %s", address))

	err = server.Serve(listener)

	if err != nil {
		log.Fatalf("ошибка запуска grpc сервера %v", err)
	}
}
