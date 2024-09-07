package config

import "os"

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Driver  string
	Schema  string
}

func LoadConfig() *Config {
	return &Config{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Driver:   os.Getenv("DB_DRIVER"),
			Schema:   os.Getenv("DB_SCHEMA"),
		},
	}
}
