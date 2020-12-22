package config

import (
	"encoding/json"
	"fmt"
	"go-example/log"
	"strconv"

	"github.com/spf13/viper"
)

var runtimeViper = viper.New()

func init() {
	log.Debug("INIT CONFIG")
}

// Config struct
type Config struct {
	Port     int
	Database struct {
		Type string
		URL  string
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
func AllConf() (*Config, error) {
	c := &Config{}
	err := runtimeViper.Unmarshal(c)
	return c, err
}

// Viper instance
func Viper() *viper.Viper {
	return runtimeViper
}

// Get config value
func Get(key string) string {
	val := runtimeViper.Get(key)
	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	case float64:
		return fmt.Sprintf("%f", val.(float64))
	default:
		return ""
	}
}
