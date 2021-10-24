package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APIUrl          string `mapstructure:"API_BASE_URL"`
	HASH            string `mapstructure:"HASH"`
	StoreName       string `mapstructure:"STORE_NAME"`
	BucketName      string `mapstructure:"BUCKET_NAME"`
	ServiceURL      string `mapstructure:"SERVICE_URL"`
	ServiceProtocol string `mapstructure:"SERVICE_PROTOCOL"`
	MongoUri        string `mapstructure:"MONGO_URI"`
	DatabaseName    string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig() (config Config) {
	viper.AddConfigPath("./util")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Could not load Config", err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}
