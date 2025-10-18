package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Daemon   DaemonConfig   `mapstructure:"daemon"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type DaemonConfig struct {
	Port int `mapstructure:"port"`
}

var Cfg Config

func LoadConfig() error {
	v := viper.New()
	v.SetDefault("log.level", "info")
	v.SetDefault("database.path", "/var/tmp/scheduleX/jobs.db")
	v.SetDefault("daemon.port", 8123)

	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AddConfigPath("/etc/scheduleX")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix("SCHEDULEX")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}
