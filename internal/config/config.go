package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName   string         `envconfig:"service_name" required:"true" default:"chasingpaper"`
	BindAddress   string         `envconfig:"bind_address" required:"true" default:":13702"`
	EnableTracing bool           `envconfig:"enable_tracing" default:"false"`
	EnableMetrics bool           `envconfig:"enable_metrics" default:"false"`
	AuthMethod    string         `envconfig:"auth_method" default:"none"` // oauth2, basic, none
	Postgres      PostgresConfig `envconfig:"postgres"`
	Kafka         KafkaConfig    `envconfig:"kafka"`
}

type PostgresConfig struct {
	DSN          string `envconfig:"dsn" required:"true"`
	MaxOpenConns int    `envconfig:"max_open_conns" default:"10"`
	MaxIdleConns int    `envconfig:"max_idle_conns" default:"2"`
}

type KafkaConfig struct {
	Brokers []string `envconfig:"brokers"`
}

func LoadConfig(envFiles ...string) *Config {
	if err := godotenv.Load(envFiles...); err != nil {
		log.Fatalf("no .env file found or error loading .env file: %v", err)
	}

	var cfg Config
	if err := envconfig.Process("CHASING_PAPER", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
