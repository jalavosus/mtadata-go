package gtfsapi

import (
	"log"

	"github.com/jalavosus/mtadata/internal/env"

	_ "github.com/joho/godotenv/autoload"
)

const (
	mtaApiKeyEnv string = "MTA_API_KEY"
	apiKeyHeader string = "x-api-key"
)

func mtaApiKey() string {
	apiKey := env.StringFromEnv(mtaApiKeyEnv, "")
	if apiKey == "" {
		log.Fatalf("%[1]s not set in env\n", mtaApiKeyEnv)
	}

	return apiKey
}
