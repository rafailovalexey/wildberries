package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

type InterfaceDatabase interface {
	Initialize()
	GetPool() *pgxpool.Pool
}

type database struct {
	credentials string
}

var _ InterfaceDatabase = (*database)(nil)

func NewDatabase() *database {
	db := &database{}

	db.Initialize()

	return db
}

func (d *database) Initialize() {
	username := os.Getenv("POSTGRESQL_USERNAME")

	if username == "" {
		log.Panicf("specify the database user")
	}

	password := os.Getenv("POSTGRESQL_PASSWORD")

	if password == "" {
		log.Panicf("specify the database user password")
	}

	db := os.Getenv("POSTGRESQL_DATABASE")

	if db == "" {
		log.Panicf("indicate the name of the database")
	}

	hostname := os.Getenv("POSTGRESQL_HOSTNAME")

	if hostname == "" {
		log.Panicf("specify the database hostname")
	}

	port := os.Getenv("POSTGRESQL_PORT")

	if port == "" {
		log.Panicf("specify the database port")
	}

	sslmode := os.Getenv("POSTGRESQL_SSLMODE")

	if sslmode == "" {
		log.Panicf("specify the ssl mode of the database")
	}

	d.credentials = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", username, password, db, hostname, port, sslmode)
}

func (d *database) GetPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), d.credentials)

	if err != nil {
		log.Panicf("error %v\n", err)
	}

	return pool
}
