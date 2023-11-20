package cron_scheduler

import (
	"fmt"
	"github.com/emptyhopes/employees_cron/internal/service"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(service service.InterfaceEmployeeService) {
	c := cron.New()

	_, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("This job runs every minute.")
		service.UpdateEmployeeWithoutConfirmation()
	})

	if err != nil {
		log.Fatalf("произошла ошибка при выполнение cron %v", err)

		return
	}

	c.Start()
	defer c.Stop()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT)
	<-exit
}
