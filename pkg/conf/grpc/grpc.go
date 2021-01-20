package yidun

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

type grpcConfig struct {
	GrpcServerPort string `mapstructure:"grpc_server_port"`
}

type config struct {
	Gc grpcConfig `mapstructure:"grpc"`
}

// Config loads the config files
func Config() *config {
	once.Do(func() {
		c = &config{}
		filename := fmt.Sprintf("configs/%v/grpc.toml", envconf.CurrentEnvironment())

		fmt.Printf("GRPC_CONFIG_FILE: %v\n", filename)

		viper.SetConfigType("toml")
		viper.SetConfigFile(filename)

		viper.ReadInConfig()
		viper.Unmarshal(c)
	})

	return c
}
