package main

import (
	"github.com/emptyhopes/employees/cmd/application"
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := application.NewApplication(ctx)

	if err != nil {
		log.Fatalf("произошла ошибка при запуске приложения %v", err)
	}

	app.Run()
}
