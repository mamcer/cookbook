package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	APIPort string `json:"api_port" env:"API_PORT"`
	WebPort string `json:"web_port" env:"WEB_PORT"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	DriverName     string `json:"driver_name" env:"DB_DRIVER_NAME"`
	DataSourceName string `json:"data_source_name" env:"DB_DATA_SOURCE_NAME"`
	MaxOpenConns   int    `json:"max_open_conns" env:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns   int    `json:"max_idle_conns" env:"DB_MAX_IDLE_CONNS"`
}

// Load loads configuration from file and environment variables
func Load(configPath string) (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			APIPort: "5001",
			WebPort: "5000",
		},
		Database: DatabaseConfig{
			DriverName:     "mysql",
			DataSourceName: "root:root@tcp(localhost:3366)/cookbook",
			MaxOpenConns:   25,
			MaxIdleConns:   5,
		},
	}

	// Load from file if it exists
	if configPath != "" {
		if err := loadFromFile(configPath, config); err != nil {
			return nil, fmt.Errorf("failed to load config from file: %w", err)
		}
	}

	// Override with environment variables
	loadFromEnv(config)

	return config, nil
}

// loadFromFile loads configuration from a JSON file
func loadFromFile(configPath string, config *Config) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(config)
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(config *Config) {
	if port := os.Getenv("API_PORT"); port != "" {
		config.Server.APIPort = port
	}
	if port := os.Getenv("WEB_PORT"); port != "" {
		config.Server.WebPort = port
	}
	if driver := os.Getenv("DB_DRIVER_NAME"); driver != "" {
		config.Database.DriverName = driver
	}
	if dsn := os.Getenv("DB_DATA_SOURCE_NAME"); dsn != "" {
		config.Database.DataSourceName = dsn
	}
	if maxOpen := os.Getenv("DB_MAX_OPEN_CONNS"); maxOpen != "" {
		if val, err := strconv.Atoi(maxOpen); err == nil {
			config.Database.MaxOpenConns = val
		}
	}
	if maxIdle := os.Getenv("DB_MAX_IDLE_CONNS"); maxIdle != "" {
		if val, err := strconv.Atoi(maxIdle); err == nil {
			config.Database.MaxIdleConns = val
		}
	}
} 