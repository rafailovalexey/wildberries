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
		log.Fatalf("укажите пользователя базы данных")
	}

	password := os.Getenv("POSTGRESQL_PASSWORD")

	if password == "" {
		log.Fatalf("укажите пользователя базы данных")
	}

	db := os.Getenv("POSTGRESQL_DATABASE")

	if db == "" {
		log.Fatalf("укажите название базы данных")
	}

	hostname := os.Getenv("POSTGRESQL_HOSTNAME")

	if hostname == "" {
		log.Fatalf("укажите имя хоста базы данных")
	}

	port := os.Getenv("POSTGRESQL_PORT")

	if port == "" {
		log.Fatalf("укажите порт базы данных")
	}

	sslmode := os.Getenv("POSTGRESQL_SSLMODE")

	if sslmode == "" {
		log.Fatalf("укажите ssl mode базы данных")
	}

	d.credentials = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", username, password, db, hostname, port, sslmode)
}

func (d *database) GetPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), d.credentials)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	return pool
}