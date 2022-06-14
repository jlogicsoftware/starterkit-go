package config

import (
	"os"
	"strconv"

	"github.com/jlogicsoftware/starterkit/internal/pkg/logger"
)

type config struct {
	Env  string
	Port int
	// Listen      struct {
	// 	Type   string `env:"LISTEN_TYPE" env-default:"port"`
	// 	BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
	// 	Port   string `env:"PORT" env-default:"3000"`
	// }
	// AppConfig struct {
	// 	LogLevel  string `env:"LOG_LEVEL" env-default:"info"`
	// 	AdminUser struct {
	// 		Email    string `env:"ADMIN_EMAIL" env-required:"true" env-default:"admin"`
	// 		Password string `env:"ADMIN_PASSWORD" env-required:"true" env-default:"admin"`
	// 	}
	// }
}

var instance *config

func init() {
	logger.Info().Msg("Loading config...")
	// Initialize default values
	instance = &config{
		Env:  "production",
		Port: 3000,
	}
	// Overload environment variables
	if e := os.Getenv("ENV"); e != "" {
		instance.Env = e
	}
	if p := os.Getenv("PORT"); p != "" {
		instance.Port, _ = strconv.Atoi(p)
	}
}

func Get() *config {
	return instance
}
