package api

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func ConnectDb(envFilename string) (*pq.Connector, error) {
	err := godotenv.Load(envFilename)
	if err != nil {
		return nil, FAILED_TO_LOAD_ENV
	}
	port, _ := strconv.Atoi(os.Getenv("Port"))
	User := os.Getenv("User")
	Password := os.Getenv("Password")
	Database := os.Getenv("Database")
	Host := os.Getenv("Host")
	cfg := pq.Config{
		Host:     Host,
		User:     User,
		Port:     uint16(port),
		Password: Password,
		Database: Database,
	}
	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (a *Application) RunDbMigration(db *sql.DB) error {
	dbm, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	
	
}
