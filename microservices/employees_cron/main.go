package main

import (
	"github.com/emptyhopes/employees_cron/cmd/cron"
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctx := context.Background()

	c, err := cron.NewCron(ctx)

	if err != nil {
		log.Panicf("произошла ошибка при запуске приложения %v", err)
	}

	c.Run()
}
