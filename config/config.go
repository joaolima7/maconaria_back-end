package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBName          string `mapstructure:"DB_NAME"`
	DBSSLMode       string `mapstructure:"DB_SSLMODE"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	JWTExpiresInMin int    `mapstructure:"JWT_EXPIRES_IN_MIN"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erro ao ler config: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

func (c *Config) GetJWTDuration() time.Duration {
	return time.Duration(c.JWTExpiresInMin) * time.Minute
}
