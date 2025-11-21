package config

import (
	"fmt"
	"os"
	"strconv"
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
	AutoMigrate     bool   `mapstructure:"AUTO_MIGRATE"`
	FTPHost         string `mapstructure:"FTP_HOST"`
	FTPPort         string `mapstructure:"FTP_PORT"`
	FTPUser         string `mapstructure:"FTP_USER"`
	FTPPassword     string `mapstructure:"FTP_PASSWORD"`
	FTPBasePath     string `mapstructure:"FTP_BASE_PATH"`
	FTPBaseURL      string `mapstructure:"FTP_BASE_URL"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if path != "" {
		if _, err := os.Stat(path); err == nil {
			viper.SetConfigFile(path)
			if err := viper.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
					return nil, fmt.Errorf("erro ao ler config: %w", err)
				}
			}
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal config: %w", err)
	}

	if autoMigrateStr := os.Getenv("AUTO_MIGRATE"); autoMigrateStr != "" {
		cfg.AutoMigrate, _ = strconv.ParseBool(autoMigrateStr)
	}

	return &cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=10s&readTimeout=30s&writeTimeout=30s&loc=Local",
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
