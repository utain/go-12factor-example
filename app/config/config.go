package config

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var runtimeViper = viper.New()

func init() {
	runtimeViper.SetDefault("port", 5000)
	runtimeViper.SetConfigType("yaml")
	runtimeViper.SetConfigName("default")
	runtimeViper.AddConfigPath("/etc/go-example")
	runtimeViper.AddConfigPath("$HOME/.go-example")
	runtimeViper.AddConfigPath("./config")
	runtimeViper.BindEnv("port")
	runtimeViper.BindEnv("logging.path")
	runtimeViper.BindEnv("database.url")
	runtimeViper.BindEnv("database.type")
	runtimeViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	runtimeViper.AutomaticEnv()
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
func AllConf() (Config, error) {
	var c Config
	err := runtimeViper.Unmarshal(&c)
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

// Read config with custom path
func Read(configPaths ...string) {
	for _, confPath := range configPaths {
		if confPath != "" {
			runtimeViper.AddConfigPath(confPath)
		}
	}
	err := runtimeViper.ReadInConfig() // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}
}
