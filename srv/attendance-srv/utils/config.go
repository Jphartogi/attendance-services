package utils

import (
	"github.com/spf13/viper"
)

//LoadConfig ...
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app.config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
