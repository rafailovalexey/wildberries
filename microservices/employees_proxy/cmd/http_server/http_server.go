package http_server

import (
	"fmt"
	"github.com/emptyhopes/employees_proxy/cmd/http_server/interceptor"
	"github.com/emptyhopes/employees_proxy/cmd/http_server/middleware"
	"github.com/emptyhopes/employees_proxy/internal/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Run(employeeApi api.InterfaceEmployeeApi) {
	router := mux.NewRouter()

	middlewares := middleware.ChainMiddleware(
		interceptor.LoggingInterceptor,
		middleware.CorsMiddleware,
		middleware.AuthenticationMiddleware,
	)

	router.Use(middlewares)

	router.NotFoundHandler = http.HandlerFunc(employeeApi.NotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(employeeApi.MethodNotAllowed)

	router.HandleFunc("/v1/employees/{id:[a-zA-Z0-9-]+}", employeeApi.GetEmployeeById).Methods("GET")
	router.HandleFunc("/v1/employees", employeeApi.CreateEmployee).Methods("POST")

	hostname := os.Getenv("HOSTNAME")

	port := os.Getenv("PORT")

	if port == "" {
		log.Panicf("specify the port")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Printf("http server starts at address %s\n", address)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Panicf("error when starting the grpc server %v\n", err)
	}
}
