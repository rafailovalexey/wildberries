package main

import (
	"context"
	"github.com/emptyhopes/employees_proxy/cmd/application"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := application.NewApplication(ctx)

	if err != nil {
		log.Panicf("произошла ошибка при инициализации %v", err)
	}

	app.Run()
}
