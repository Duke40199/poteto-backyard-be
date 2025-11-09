package config

import (
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(os.Getenv("CNF_DIR"))
	viper.AddConfigPath(".")
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
}

func UnmarshalKey(key string, i interface{}) error {
	return viper.UnmarshalKey(key, i, func(config *mapstructure.DecoderConfig) {
		config.TagName = "config"
	})
}

func Unmarshal(i interface{}) error {
	return viper.GetViper().Unmarshal(i, func(config *mapstructure.DecoderConfig) {
		config.TagName = "config"
	})
}

func LoadConfig(i interface{}) error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return Unmarshal(i)
}
