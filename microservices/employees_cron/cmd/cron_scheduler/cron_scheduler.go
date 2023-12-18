package cron_scheduler

import (
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
		service.UpdateEmployeeWithoutConfirmation()
	})

	if err != nil {
		log.Panicf("n error occurred while executing cron %v", err)

		return
	}

	log.Printf("cron started\n")

	c.Start()
	defer c.Stop()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT)
	<-exit
}
