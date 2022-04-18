package viperadapter

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/utain/go/example/internal/errs"

	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var config Config

// Config struct
type Config struct {
	Database struct {
		URL  string
		Pool struct {
			MaxOpen int
			MaxIdle int
		}
	}
}

func (d Config) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// Parse get all config support in app
func Parse() error {
	viperInstance.SetConfigName("config")
	viperInstance.AddConfigPath("./config")
	viperInstance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperInstance.AutomaticEnv()
	if err := viperInstance.ReadInConfig(); err != nil {
		log.Fatalf("Load config from file [%s]: %v", viperInstance.ConfigFileUsed(), err)
	}
	if err := viperInstance.Unmarshal(&config); err != nil {
		return errs.ErrInvalidConfig.With("err", err)
	}
	return nil
}

// Config values
func V() Config {
	return config
}
