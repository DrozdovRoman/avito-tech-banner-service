package configuration

type PostgresConfiguration struct {
	DBName   string `json:"database" required:"true" default:"postgres"`
	Host     string `json:"host" required:"true" default:"localhost"`
	Port     int    `json:"port" required:"true" default:"5432"`
	User     string `json:"user" required:"true" default:"user"`
	Password string `json:"password" required:"true" default:"postgres"`
}
