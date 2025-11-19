package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	Name     string
	Category string
	Protocol string
	Host     string
	Path     string
	Timeout  int
}

type Config struct {
	Port     string
	Services []ServiceConfig
}

func LoadConfig() Config {
	godotenv.Load()

	port := getEnv("PORT", "8081")

	totalServices, err := strconv.Atoi(getEnv("SERVICES", "0"))
	if err != nil {
		log.Fatalf("SERVICES debe ser un número entero")
	}

	services := make([]ServiceConfig, 0)

	for i := 1; i <= totalServices; i++ {
		prefix := fmt.Sprintf("SERVICE_%d_", i)

		svc := ServiceConfig{
			Name:     getEnv(prefix+"NAME", fmt.Sprintf("Service %d", i)),
			Category: getEnv(prefix+"CATEGORY", "General"),
			Protocol: getEnv(prefix+"PROTOCOL", "https"),
			Host:     getEnv(prefix+"HOST", ""),
			Path:     getEnv(prefix+"PATH", "/"),
			Timeout:  getEnvInt(prefix+"TIMEOUT", 5),
		}

		if svc.Host == "" {
			log.Fatalf("%sHOST no puede estar vacío", prefix)
		}

		services = append(services, svc)
	}

	return Config{
		Port:     port,
		Services: services,
	}
}

// Helpers

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
