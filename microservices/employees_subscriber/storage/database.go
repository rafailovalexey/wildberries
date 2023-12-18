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
		log.Panicf("укажите пользователя базы данных")
	}

	password := os.Getenv("POSTGRESQL_PASSWORD")

	if password == "" {
		log.Panicf("укажите пользователя базы данных")
	}

	db := os.Getenv("POSTGRESQL_DATABASE")

	if db == "" {
		log.Panicf("укажите название базы данных")
	}

	hostname := os.Getenv("POSTGRESQL_HOSTNAME")

	if hostname == "" {
		log.Panicf("укажите имя хоста базы данных")
	}

	port := os.Getenv("POSTGRESQL_PORT")

	if port == "" {
		log.Panicf("укажите порт базы данных")
	}

	sslmode := os.Getenv("POSTGRESQL_SSLMODE")

	if sslmode == "" {
		log.Panicf("укажите ssl mode базы данных")
	}

	d.credentials = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", username, password, db, hostname, port, sslmode)
}

func (d *database) GetPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), d.credentials)

	if err != nil {
		log.Panicf("ошибка %v\n", err)
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
		log.Panicf("ошибка создания таблицы: %v\n", err)
	}
}
