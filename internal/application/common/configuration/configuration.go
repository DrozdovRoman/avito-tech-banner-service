package configuration

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Configuration struct {
	Postgres PostgresConfiguration
	HTTP     HTTPConfiguration
}

type PostgresConfiguration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type HTTPConfiguration struct {
	IP   string `json:"ip" required:"true" default:"0.0.0.0"`
	Port int    `json:"port" required:"true" default:"8000"`
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Overload(filename)
	} else {
		err = godotenv.Load()
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadConfiguration(filename string) (*Configuration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("template", config); err != nil {
		return nil, err
	}

	return config, nil
}
