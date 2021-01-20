package elasticsearch

import (
	"fmt"
	"sync"
	"github.com/spf13/viper"
	envconf "github.com/mwy001/goland/pkg/conf/env"
)

var (
	c    *config
	once sync.Once
)

type esConfig struct {
	Address string `mapstructure:"address"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
}

type config struct {
	Es esConfig `mapstructure:"elasticsearch"`
}

// Config loads the config files
func Config() *config {
	once.Do(func() {
		c = &config{}
		filename := fmt.Sprintf("configs/%v/elasticsearch.toml", envconf.CurrentEnvironment())
		fmt.Printf("ELASTICSEARCH_CONFIG_FILE: %v\n", filename)

		viper.SetConfigType("toml")
		viper.SetConfigFile(filename)

		viper.ReadInConfig()
		viper.Unmarshal(c)
	})

	return c
}
