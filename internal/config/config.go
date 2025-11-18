package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	Services       []string
	TimeoutSeconds int
}

func LoadConfig() Config {
	godotenv.Load()

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_SECONDS"))
	if err != nil {
		timeout = 5
	}

	return Config{
		Port:           os.Getenv("PORT"),
		Services:       strings.Split(os.Getenv("SERVICES"), ","),
		TimeoutSeconds: timeout,
	}
}
