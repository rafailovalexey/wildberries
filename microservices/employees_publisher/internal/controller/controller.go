package controller

import "github.com/nats-io/stan.go"

type InterfaceEmployeeController interface {
	PublishEmployee(stan.Conn, string)
}
