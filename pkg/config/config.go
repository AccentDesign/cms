package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net"
	"net/url"
	"strings"
)

type (
	SslMode string
)

//goland:noinspection GoUnusedConst
const (
	SslModeDisable SslMode = "disable"
	SslModeAllow   SslMode = "allow"
	SslModePrefer  SslMode = "prefer"
	SslModeRequire SslMode = "require"
)

// Config represents the top-level configuration structure.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Security SecurityConfig `mapstructure:"security"`
}

// FromPath creates and validates a new Config from a .toml file.
// If environment variables are set, they will override the values from the .toml file.
// The environment variables must be prefixed with the
// name of the configuration structure in uppercase, and the keys must be separated
// by underscores. For example, to override the `port` value in the `server` structure,
// the environment variable must be `SERVER_PORT`.
func FromPath(path string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}

// ServerConfig represents the server configuration.
type ServerConfig struct {
	Port  uint16 `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
	Url   string `mapstructure:"url"`
}

// DatabaseConfig represents the database configuration.
type DatabaseConfig struct {
	Host     string  `mapstructure:"host"`
	Port     uint16  `mapstructure:"port"`
	User     string  `mapstructure:"user"`
	Password string  `mapstructure:"password"`
	Db       string  `mapstructure:"db"`
	SslMode  SslMode `mapstructure:"ssl_mode"`
}

// SecurityConfig represents the security configuration.
type SecurityConfig struct {
	AllowedHosts          []string `mapstructure:"allowed_hosts"`
	HSTSMaxAge            int      `mapstructure:"hsts_max_age"`
	XSSProtection         string   `mapstructure:"xss_protection"`
	ContentTypeNosniff    string   `mapstructure:"content_type_nosniff"`
	XFrameOptions         string   `mapstructure:"x_frame_options"`
	ContentSecurityPolicy string   `mapstructure:"content_security_policy"`
	ReferrerPolicy        string   `mapstructure:"referrer_policy"`
}

// URL returns the database URL.
func (c DatabaseConfig) URL() *url.URL {
	query := url.Values{}
	query.Set("sslmode", string(c.SslMode))
	return &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.User, c.Password),
		Host:     net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.Port)),
		Path:     c.Db,
		RawQuery: query.Encode(),
	}
}
