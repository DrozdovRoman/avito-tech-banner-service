package configuration

type HTTPConfiguration struct {
	IP   string `json:"ip" required:"true" default:"0.0.0.0"`
	Port int    `json:"port" required:"true" default:"8000"`
}
