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
	CreateTables(*pgxpool.Pool)
}

type database struct {
	credentials string
}

var _ InterfaceDatabase = (*database)(nil)

func NewDatabase() *database {
	db := &database{}

	db.Initialize()

	pool := db.GetPool()
	defer pool.Close()

	db.CreateTables(pool)

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

func (d *database) CreateTables(pool *pgxpool.Pool) {
	createEmployeeTable(pool)
}

func createEmployeeTable(pool *pgxpool.Pool) {
	query := `
		CREATE TABLE IF NOT EXISTS employees (
			employee_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
			confirmation BOOL,
			firstname VARCHAR(255),
			lastname VARCHAR(255),
			email VARCHAR(255),
			phone_number VARCHAR(255),
			address VARCHAR(255),
			position VARCHAR(255),
			department VARCHAR(255),
			date_of_birth TIMESTAMP,
			hire_date TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := pool.Exec(context.Background(), query)

	if err != nil {
		log.Panicf("table creation error %v\n", err)
	}
}
