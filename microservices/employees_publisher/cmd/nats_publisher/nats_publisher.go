package nats_publisher

import (
	"fmt"
	"github.com/emptyhopes/employees_publisher/internal/controller"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

func Run(employeeController controller.InterfaceEmployeeController) {
	sc := connect()

	defer sc.Close()

	queue := "employees"

	fmt.Printf("подписался на очередь %s\n", queue)

	for {
		employeeController.PublishEmployee(sc, queue)

		time.Sleep(1 * time.Second)
	}
}

func connect() stan.Conn {
	url := os.Getenv("NATS_URL")

	if url == "" {
		log.Fatalf("укажите nats url")
	}

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Fatalf("укажите идентификатор кластера")
	}

	sc, err := stan.Connect(cluster, "publisher-1", stan.NatsURL(url))

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	return sc
}
