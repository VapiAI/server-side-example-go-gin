package config

import (
	"os"
)

type EnvConfig struct {
	Weather struct {
		BaseUrl string
		ApiKey  string
	}
	Openai struct {
		ApiKey string
	}
	Vapi struct {
		BaseUrl string
		ApiKey  string
	}
}

func LoadEnvConfig() EnvConfig {
	return EnvConfig{
		Weather: struct {
			BaseUrl string
			ApiKey  string
		}{
			BaseUrl: getEnv("WEATHER_BASE_URL", "https://api.openweathermap.org/data/2.5"),
			ApiKey:  getEnv("WEATHER_API_KEY", ""),
		},
		Openai: struct {
			ApiKey string
		}{
			ApiKey: getEnv("OPENAI_API_KEY", ""),
		},
		Vapi: struct {
			BaseUrl string
			ApiKey  string
		}{
			BaseUrl: getEnv("VAPI_BASE_URL", "https://api.vapi.ai"),
			ApiKey:  getEnv("VAPI_API_KEY", ""),
		},
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
