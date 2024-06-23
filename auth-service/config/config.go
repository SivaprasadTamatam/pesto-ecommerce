package config

import "os"

type Config struct {
	DBDriver  string
	DBSource  string
	JwtSecret string
}

func LoadConfig() Config {
	return Config{
		DBDriver:  "sqlite3",
		DBSource:  "test.db",
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
