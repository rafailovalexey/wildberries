package main

import (
	"github.com/emptyhopes/employees_subscriber/cmd/subscriber"
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctx := context.Background()

	sub, err := subscriber.NewSubscriber(ctx)

	if err != nil {
		log.Panicf("an error occurred while starting the application %v\n", err)
	}

	sub.Run()
}
