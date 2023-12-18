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
		log.Panicf("укажите порт")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Printf("%s\n", fmt.Sprintf("http сервер запускается по адресу %s", address))

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Panicf("ошибка при запуске сервера: %v\n", err)
	}
}
