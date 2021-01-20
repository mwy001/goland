package mongodb

import (
	"sync"
	"fmt"
	"github.com/spf13/viper"
	envconf "github.com/mwy001/goland/pkg/conf/env"
)

var (
	c    *config
	once sync.Once
)

type mongoConfig struct {
	Address string `mapstructure:"address"`
	DB      string `mapstructure:"db"`
}

type config struct {
	Mongo mongoConfig `mapstructure:"mongodb"`
}

// Config loads the config files
func Config() *config {
	once.Do(func() {
		c = &config{}
		filename := fmt.Sprintf("configs/%v/mongodb.toml", envconf.CurrentEnvironment())
		fmt.Printf("MONGODB_CONFIG_FILE: %v\n", filename)

		viper.SetConfigType("toml")
		viper.SetConfigFile(filename)

		viper.ReadInConfig()
		viper.Unmarshal(c)
	})

	return c
}
