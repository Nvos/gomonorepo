package infra

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Profile int

const (
	Development Profile = iota
	Production
)

type Logging struct {
	LogLevel    zapcore.Level
	HttpLogging bool
}

type Environment struct {
	Profile Profile
}

type Server struct {
	Host string
	Port int
}

type Security struct {
	JwtSigningSecret        string
	JwtTokenLifetime        int
	JwtRefreshTokenLifetime int
}

type Configuration struct {
	Environment *Environment
	Server      *Server
	Security    *Security
	Logging     *Logging
}

func (c *Configuration) IsDevelopment() bool {
	return c.Environment.Profile == Development
}

func (c *Configuration) IsProduction() bool {
	return c.Environment.Profile == Production
}

func defaultViperConfig() *viper.Viper {
	v := viper.New()
	v.SetDefault("environment.profile", Development)

	v.SetDefault("logging.logLevel", zap.InfoLevel)

	v.SetDefault("server.host", "http://localhost")
	v.SetDefault("server.port", 8080)

	v.SetDefault("security.jwtTokenLifetime", time.Minute*30)
	v.SetDefault("security.jwtRefreshTokenLifetime", time.Minute*60*24)

	return v
}

func MustNewConfiguration() *Configuration {
	v := defaultViperConfig()

	v.SetConfigFile("config.toml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to read configuration, error: %v\n", err))
	}

	var configuration *Configuration
	err = v.Unmarshal(&configuration)
	if err != nil {
		panic(fmt.Sprintf("Failed to un marshall configuration, error: %v\n", err))
	}

	return configuration
}
