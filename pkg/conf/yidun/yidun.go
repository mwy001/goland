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

type yidunConfig struct {
	SecretID  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`

	BusIDImage            string `mapstructure:"busid_image"`
	ImageCheckURL         string `mapstructure:"image_check_url"`
	ImageCheckVersion     string `mapstructure:"image_check_version"`
	ImageQueryTaskURL     string `mapstructure:"image_query_task_url"`
	ImageQueryTaskVersion string `mapstructure:"image_query_task_version"`

	BusIDVideo            string `mapstructure:"busid_video"`
	VideoSubmitURL        string `mapstructure:"video_submit_url"`
	VideoSubmitVersion    string `mapstructure:"video_submit_version"`
	VideoQueryTaskURL     string `mapstructure:"video_query_task_url"`
	VideoQueryTaskVersion string `mapstructure:"video_query_task_version"`

	SpamRatingThreshold float64 `mapstructure:"spam_rating_threshold"`
}

type config struct {
	YidunConfig yidunConfig `mapstructure:"yidun"`
}

// Config loads the config files
func Config() *config {
	once.Do(func() {
		c = &config{}
		filename := fmt.Sprintf("configs/%v/yidun.toml", envconf.CurrentEnvironment())
		fmt.Printf("YIDUN_CONFIG_FILE: %v\n", filename)

		viper.SetConfigType("toml")
		viper.SetConfigFile(filename)

		viper.ReadInConfig()
		viper.Unmarshal(c)
	})

	return c
}
