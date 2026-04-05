package config

type PalPayConfig struct {
	ClientID     string `envconfig:"client_id"`
	ClientSecret string `envconfig:"client_secret"`
}

type StripeConfig struct {
	SecretKey string `envconfig:"secret_key"`
}
