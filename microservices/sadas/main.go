package main

import (
	"github.com/emptyhopes/employees_publisher/cmd/publisher"
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctx := context.Background()

	pub, err := publisher.NewPublisher(ctx)

	if err != nil {
		log.Fatalf("произошла ошибка при запуске приложения %v", err)
	}

	pub.Run()
}
