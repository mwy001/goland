package obs

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

type obsConfig struct {
	AK string `mapstructure:"ak"`
	SK string `mapstructure:"sk"`

	Bucket   string `mapstructure:"bucket"`
	BucketG2 string `mapstructure:"bucket_g2"`

	Endpoint   string `mapstructure:"endpoint"`
	EndpointG2 string `mapstructure:"endpoint_g2"`
}

type config struct {
	Obs obsConfig `mapstructure:"obs"`
}

// Config loads the config files
func Config() *config {
	once.Do(func() {
		c = &config{}
		filename := fmt.Sprintf("configs/%v/obs.toml", envconf.CurrentEnvironment())
		fmt.Printf("OBS_CONFIG_FILE: %v\n", filename)

		viper.SetConfigType("toml")
		viper.SetConfigFile(filename)

		viper.ReadInConfig()
		viper.Unmarshal(c)
	})

	return c
}
