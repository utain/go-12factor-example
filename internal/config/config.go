package config

import (
	"encoding/json"
	"go-example/log"

	"github.com/spf13/viper"
)

var defaultViper = viper.New()
var defaultConfig Config

func init() {
	log.Debug("INIT CONFIG")
}

// Config struct
type Config struct {
	Port     int
	Database struct {
		URL string
	}
	Logging struct {
		Path string
	}
}

func (d Config) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// AllConf get all config support in app
func AllConf() Config {
	err := defaultViper.Unmarshal(&defaultConfig)
	if err != nil {
		log.Fatal("Fail to read configuration")
	}
	log.Debug("defaultConfig:", defaultConfig)
	return defaultConfig
}

// Viper instance
func Viper() *viper.Viper {
	return defaultViper
}
