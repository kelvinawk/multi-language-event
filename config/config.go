package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	Env    string `mapstructure:"ENV"`
	Db_dsn string `mapstructure:"DB_DSN"`
}

func Load() *Config {
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("PORT", 8080)

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	return &Config{
		Port:   viper.GetString("PORT"),
		Env:    viper.GetString("ENV"),
		Db_dsn: viper.GetString("DB_DSN"),
	}
}
