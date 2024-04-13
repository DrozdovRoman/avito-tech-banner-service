package configuration

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Configuration struct {
	ApiExternalURL string `json:"api_external_url" split_words:"true" required:"true"`
	SecretKey      string `json:"secret_key" split_words:"true" required:"true"`

	Postgres PostgresConfiguration `json:"postgres" required:"true"`
	Cache    CacheConfiguration    `json:"cache" required:"true"`
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
