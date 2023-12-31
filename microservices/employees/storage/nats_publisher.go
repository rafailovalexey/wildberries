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
		log.Panicf("specify nats url")
	}

	p.url = url

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Panicf("specify the cluster id")
	}

	p.cluster = cluster
}

func (p *publisher) GetConnect() stan.Conn {
	sc, err := stan.Connect(p.cluster, "publisher-1", stan.NatsURL(p.url))

	if err != nil {
		log.Panicf("error %v\n", err)
	}

	return sc
}
