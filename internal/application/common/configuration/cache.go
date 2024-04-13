package configuration

import "time"

type CacheConfiguration struct {
	Expiration time.Duration `json:"expiration" required:"true" default:"5m"`
	Cleanup    time.Duration `json:"cleanup" required:"true" default:"10m"`
}
