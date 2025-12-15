package config

import (
	"os"
)

func Load() *Config {
	return &Config{
		LinkedInEmail:    os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword: os.Getenv("LINKEDIN_PASSWORD"),

		LogLevel:        getEnv("LOG_LEVEL", "info"),
		BrowserHeadless: getEnv("BROWSER_HEADLESS", "true") == "true",
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
