package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL       string
	DisableSend  bool
	DisablePrint bool
	Concurrency  int
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or failed to load, proceeding with environment variables")
	}

	concurrency := 1
	if val := os.Getenv("CONCURRENCY"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 {
			concurrency = parsed
		}
	}

	AppConfig = Config{
		APIURL:       os.Getenv("API_URL"),
		DisableSend:  os.Getenv("DISABLE_SEND") == "1" || os.Getenv("DISABLE_SEND") == "true",
		DisablePrint: os.Getenv("DISABLE_PRINT") == "1" || os.Getenv("DISABLE_PRINT") == "true",
		Concurrency:  concurrency,
	}
}
