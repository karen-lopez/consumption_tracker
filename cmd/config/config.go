package config

import "fmt"

//goland:noinspection SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection
type Config struct {
	ServerPort        string `mapstructure:"SERVER_PORT"`
	DBURL             string `mapstructure:"DB_URL"`
	AddressServiceURL string `mapstructure:"ADDRESS_SERVICE_URL"`
	APIToken          string `mapstructure:"API_TOKEN"`
	PostgresUser      string `mapstructure:"POSTGRES_USER"`
	PostgresPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB        string `mapstructure:"POSTGRES_DB"`
}

func LoadConfig() (*Config, error) {
	cfg, err := LoadEnv()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return cfg, nil
}
