package config

import (
	"fmt"

	"github.com/hmoragrega/f3-payments/pkg/persistence"
	"github.com/spf13/viper"
)

const (
	f3EnvVarsPrefix           = "f3_api"
	f3APIDefaultIP            = "127.0.0.1"
	f3APIDefaultPort          = "8080"
	f3APIDefaultLogFormat     = "method=${method}, uri=${uri}, status=${status}\n"
	f3APIDefaultMongoServer   = "127.0.0.1:27017"
	f3APIDefaultMongoAuthDB   = "admin"
	f3APIDefaultMongoUser     = "root"
	f3APIDefaultMongoPass     = "demo"
	f3APIDefaultMongoDatabase = "f3api"
)

// Config represents the application configurable values
type Config struct {
	ServerIP   string
	ServerPort string
	LogFormat  string
	Database   string
	Mongo      persistence.MongoConfig
}

// NewConfig allows to read the env variables as config values
func NewConfig() *Config {
	viper.SetEnvPrefix(f3EnvVarsPrefix)

	return &Config{
		ServerIP:   config("ip", f3APIDefaultIP),
		ServerPort: config("port", f3APIDefaultPort),
		LogFormat:  config("log_format", f3APIDefaultLogFormat),
		Mongo: persistence.MongoConfig{
			Address:  config("mongo_server", f3APIDefaultMongoServer),
			AuthDB:   config("mongo_auth_db", f3APIDefaultMongoAuthDB),
			User:     config("mongo_user", f3APIDefaultMongoUser),
			Pass:     config("mongo_pass", f3APIDefaultMongoPass),
			Database: config("mongo_database", f3APIDefaultMongoDatabase),
		},
	}
}

// GetServerAddress returns the server address
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.ServerIP, c.ServerPort)
}

// GetAPIEndpoint returns the api endpoint
func (c *Config) GetAPIEndpoint() string {
	return fmt.Sprintf("http://%s", c.GetServerAddress())
}

func config(key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	viper.BindEnv(key)

	return viper.GetString(key)
}
