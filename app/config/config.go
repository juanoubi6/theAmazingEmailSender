package config

import "os"

type Config struct {
	ENV  string
	PORT string

	NATS_URL string

	SENDGRID_KEY_ID string

	WORKER_AMOUNT string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		config := newConfig()
		instance = &config
	}
	return instance
}

func newConfig() Config {
	return Config{
		ENV:  GetEnv("ENV", "develop"),
		PORT: GetEnv("PORT", "5002"),

		NATS_URL: GetEnv("NATS_URL", "0.0.0.0:4222"),

		SENDGRID_KEY_ID: GetEnv("SENDGRID_KEY_ID", ""),

		WORKER_AMOUNT: GetEnv("WORKER_AMOUNT", "3"),
	}
}

func GetEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}
