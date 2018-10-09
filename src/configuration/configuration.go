package configuration

import (
	"time"

	"github.com/spf13/viper"
)

const (
	Development = "development"
	Production  = "production"
	Testing     = "testing"
)

var config = new(configuration)

func Init() error {
	v := viper.New()
	setDefaults(v)

	if err := v.Unmarshal(config); err != nil {
		return err
	}

	switch config.Environment {
	case Production, Testing:
		// Do nothing
	default:
		config.Environment = Development
	}

	return nil
}

func setDefaults(v *viper.Viper) {
	v.SetEnvPrefix("URL_SHORTENER")

	for env, def := range map[string]interface{}{
		"environment":          Development,
		"show_config":          false,
		"bind":                 "0.0.0.0:5000",
		"http_host":            "127.0.0.1:5000",
		"http_scheme":          "http",
		"graceful_timeout":     5 * time.Second,
		"cors_allowed_origins": "http://localhost:5000,http://127.0.0.1:5000",
	} {
		v.BindEnv(env)
		v.SetDefault(env, def)
	}
}

type configuration struct {
	Environment        string        `mapstructure:"environment"`
	Bind               string        `mapstructure:"bind"`
	ShowConfig         bool          `mapstructure:"show_config"`
	Host               string        `mapstructure:"host"`
	Scheme             string        `mapstructure:"http_scheme"`
	CORSAllowedOrigins []string      `mapstructure:"cors_allowed_origins"`
	GracefulTimeout    time.Duration `mapstructure:"graceful_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	IdleTimeout        time.Duration `mapstructure:"idle_timeout"`
}

func Environment() string {
	return config.Environment
}

func ShowConfig() bool {
	return config.ShowConfig
}

func ReadTimeout() time.Duration {
	return config.ReadTimeout
}

func WriteTimeout() time.Duration {
	return config.WriteTimeout
}

func IdleTimeout() time.Duration {
	return config.IdleTimeout
}

func Bind() string {
	return config.Bind
}

func GracefulTimeout() time.Duration {
	return config.GracefulTimeout
}

func CORSAllowedOrigins() []string {
	return config.CORSAllowedOrigins
}

func IsDevelopment() bool {
	return config.Environment == Development
}

func IsProduction() bool {
	return config.Environment == Production
}

func IsTesting() bool {
	return config.Environment == Testing
}
