package configuration

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Configuration struct {
	Postgres PostgresConfiguration `json:"postgres" required:"true"`
	HTTP     HTTPConfiguration     `json:"http" required:"true"`
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Overload(filename)
	} else {
		err = godotenv.Load()
		// handle if .env file does not exist, this is OK
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
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
