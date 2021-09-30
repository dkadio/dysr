package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APIUrl          string `mapstructure:"API_BASE_URL"`
	HASH            string `mapstructure:"HASH"`
	StoreName       string `mapstructure:"STORE_NAME"`
	BucketName      string `mapstructure:"BUCKET_NAME"`
	ServiceURL      string `mapstructure:"SERVICE_URL"`
	ServiceProtocol string `mapstructure:"SERVICE_PROTOCOL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}
