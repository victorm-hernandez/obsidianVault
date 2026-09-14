package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	yaml "go.yaml.in/yaml/v3"
)

type WebServerConfig struct {
	Address      string
	WriteTimeout time.Duration `yaml:"write_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	IddleTimeout time.Duration `yaml:"iddle_timeout"`
}

type ObservabilityConfig struct {
	LoggingLevel     slog.Level `yaml:"logging_level"`
	EnableProfiling  bool       `yaml:"enable_profiling"`
	ProfilingAddress string     `yaml:"profiling_address"`
}

type WeatherGeneratorConfig struct {
	BufferSize      int           `yaml:"buffer_size"`
	ClientCount     int           `yaml:"client_count"`
	MaxRetryCount   int           `yaml:"max_retry_count"`
	MinRetryDelay   time.Duration `yaml:"min_retry_delay"`
	MaxRetryDelay   time.Duration `yaml:"max_retry_delay"`
	MaxRecordAge    time.Duration `yaml:"max_record_age"`
	RequestTimeout  time.Duration `yaml:"request_timeout"`
	MaxFetchPerSec  int           `yaml:"max_fetch_per_sec"`
	RandomLocAPIURL string        `yaml:"random_loc_api_url"`
	PointsAPIURL    string        `yaml:"points_api_url"`
}

type Config struct {
	WebServer        WebServerConfig
	Observability    ObservabilityConfig
	WeatherGenerator WeatherGeneratorConfig `yaml:"weather_generator"`
}

func LoadConfig() (*Config, error) {
	var config = &Config{}

	// Attempt to find the configuration file on the following paths
	// 1. Parameter passed to executable (TODO)
	// 2. Current directory
	// 3. Config folder

	env := os.Getenv("ENV")

	if env == "" {
		env = "dev"
	}

	filename := fmt.Sprintf("%v.yaml", env)

	paths := []string{
		filename,
		filepath.Join("configs", filename),
		filepath.Join("..", "..", "configs", filename),
	}

	var data []byte
	var err error

	for _, path := range paths {
		data, err = os.ReadFile(path)

		if err == nil {
			// Success
			break
		}

		absPath, _ := filepath.Abs(path)
		slog.Debug("Failed to load config from path.", slog.String("Path", absPath))
	}

	if err != nil {
		return nil, fmt.Errorf("Error Reading Configuration File. Error:%v", err)
	}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		return nil, fmt.Errorf("Error parsing configuration file. Error: %v", err.Error())
	}

	err = config.Validate()

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) Validate() error {
	wc := c.WebServer

	if wc.Address == "" {
		return fmt.Errorf("Configuration Error. WebServer Address cannot be empty.")
	}

	if wc.WriteTimeout <= 0 || wc.ReadTimeout <= 0 || wc.IddleTimeout <= 0 {
		return fmt.Errorf("Configuration Error. Invalid/Missing Web Server Timeout.")
	}

	oc := c.Observability

	if oc.EnableProfiling && oc.ProfilingAddress == "" {
		return fmt.Errorf("Configuration Error. ObservabilityConfig.ProfilingAddress cannot be empty.")
	}

	wcg := c.WeatherGenerator

	if wcg.BufferSize <= 0 || wcg.ClientCount <= 0 || wcg.MaxFetchPerSec <= 0 || wcg.MaxRetryCount <= 0 {
		// TODO: Validate the remaining settings
	}

	return nil
}
