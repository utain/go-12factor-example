package config

import (
	"encoding/json"
	"go-example/internal/log"

	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var Default Config

func init() {
	log.Debug("INIT CONFIG")
}

// Config struct
type Config struct {
	Server struct {
		Port uint
		Host string
	}
	Database struct {
		URL  string
		Pool struct {
			Max uint
		}
	}
}

func (d Config) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// Parse get all config support in app
func Parse() Config {
	if err := viperInstance.Unmarshal(&Default); err != nil {
		log.Fatal("Fail to read configuration", err)
	}
	return Default
}

// Viper instance
func Viper() *viper.Viper {
	return viperInstance
}
