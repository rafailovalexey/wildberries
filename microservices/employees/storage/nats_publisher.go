package storage

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

type InterfaceNatsPublisher interface {
	Initialize()
	GetConnect() stan.Conn
}

type publisher struct {
	url     string
	cluster string
}

var _ InterfaceNatsPublisher = (*publisher)(nil)

func NewNatsPublisher() *publisher {
	pub := &publisher{}

	pub.Initialize()

	return pub
}

func (p *publisher) Initialize() {
	url := os.Getenv("NATS_URL")

	if url == "" {
		log.Fatalf("укажите nats url")
	}

	p.url = url

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Fatalf("укажите идентификатор кластера")
	}

	p.cluster = cluster
}

func (p *publisher) GetConnect() stan.Conn {
	sc, err := stan.Connect(p.cluster, "publisher-1", stan.NatsURL(p.url))

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	return sc
}
