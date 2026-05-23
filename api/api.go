package api

import "log"

type Application struct {
	configs
}

type configs struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}
