package api

import (
	"database/sql"
	"errors"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func OpenDB(conn *pq.Connector) (*sql.DB, error) {
	db := sql.OpenDB(conn)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
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
	migra, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", dbm)
	if err != nil {
		return err
	}
	err = migra.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}
	return nil
}
