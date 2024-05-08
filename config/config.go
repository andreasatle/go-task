package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Name string
}

type TslConfig struct {
	CertFile string
	KeyFile string
}
type Config struct {
	Server ServerConfig
	Tsl TslConfig
	Database DatabaseConfig
}
// LoadConfig loads the configuration from the config file
func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 443)
	viper.SetDefault("certFile", "cert.pem")
	viper.SetDefault("keyFile", "key.pem")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	return Config{
		Server: ServerConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
		},
		Tsl: TslConfig{
			CertFile: viper.GetString("tsl.certFile"),
			KeyFile: viper.GetString("tsl.keyFile"),
		},
		Database: DatabaseConfig{
			Host: viper.GetString("database.host"),
			Port: viper.GetInt("database.port"),
			User: viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			Name: viper.GetString("database.name"),
		},
	}
}