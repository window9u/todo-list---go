package config

import (
	"io"
	"log"
)

type Config struct {
	logger *log.Logger
}

func InitConfig(w io.Writer) Config {
	return Config{logger: log.New(w, "", log.Ldate|log.Ltime)}
}
