package api

import (
	"os"
	"strconv"

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
		Host: H,
	}

}
