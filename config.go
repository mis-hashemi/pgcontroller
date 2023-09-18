package pgcontroller

import "fmt"

type Config struct {
	Username string `env:"POSTGRES_USERNAME,notEmpty";koanf:"username"`
	Password string `env:"POSTGRES_PASSWORD,notEmpty";koanf:"password"`
	Port     int    `env:"POSTGRES_PORT,notEmpty";koanf:"port"`
	Host     string `env:"POSTGRES_HOST,notEmpty";koanf:"host"`
	DBName   string `env:"POSTGRES_DBNAME,notEmpty";koanf:"dbname"`
	Driver   string `env:"POSTGRES_DRIVER";koanf:"driver"`
	Schema   string `env:"POSTGRES_SCHEMA";koanf:"schema"`
}

const Public = "public"

func (config *Config) DSN() string {
	if config.Schema == "" {
		config.Schema = Public
	}

	return fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DBName)
}
