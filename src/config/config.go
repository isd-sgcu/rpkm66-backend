package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Service struct {
	File string `mapstructure:"file"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSL      string `mapstructure:"ssl"`
}

type App struct {
	Port         int  `mapstructure:"port"`
	Debug        bool `mapstructure:"debug"`
	BaanCacheTTL int  `mapstructure:"baan_cache_ttl"`
}

type Config struct {
	Service  Service  `mapstructure:"service"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	App      App      `mapstructure:"app"`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return
}
