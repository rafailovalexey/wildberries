package http_server

import (
	"fmt"
	"github.com/emptyhopes/employees_proxy/cmd/http_server/interceptor"
	"github.com/emptyhopes/employees_proxy/cmd/http_server/middleware"
	"github.com/emptyhopes/employees_proxy/internal/api"
	"log"
	"net/http"
	"os"
)

func Run(employeeApi api.InterfaceEmployeeApi) {
	router := http.NewServeMux()

	middlewares := middleware.ChainMiddleware(
		interceptor.LoggingInterceptor,
		middleware.CorsMiddleware,
		middleware.AuthenticationMiddleware,
	)

	router.Handle("/v1/employees/", middlewares(http.HandlerFunc(employeeApi.EmployeesHandler)))

	hostname := os.Getenv("HOSTNAME")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalf("укажите порт")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Println(fmt.Sprintf("сервер запускается по адресу %s", address))

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Fatalf("ошибка при запуске сервера: %v\n", err)
	}
}
