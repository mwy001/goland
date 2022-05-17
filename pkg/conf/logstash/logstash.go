package logstash

import (
	"fmt"
	"sync"

	envconf "github.com/mwy001/goland/pkg/conf/env"
	"github.com/spf13/viper"
)

var (
	c    *config
	once sync.Once
)

type logstatshConfig struct {
	LogstashDestinationURL string `mapstructure:"logstash_destination_url"`
}

type config struct {
	Lc logstatshConfig `mapstructure:"logstash"`
}

// Config loads the config files
func Config() *config {
	once.Do(func() {
		c = &config{}
		filename := fmt.Sprintf("configs/%v/logstash.toml", envconf.CurrentEnvironment())

		fmt.Printf("LOGSTASH_CONFIG_FILE: %v\n", filename)

		viper.SetConfigType("toml")
		viper.SetConfigFile(filename)

		viper.ReadInConfig()
		viper.Unmarshal(c)
	})

	return c
}
