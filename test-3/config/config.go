package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	envPrefix = "GITS_TEST3"
)

// InitConfig : initial config
func InitConfig(configPath, serviceName string) (*ConfigStruct, error) {
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.SetConfigName(".config")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if serviceName != "" {
		envPrefix = strings.ToUpper(serviceName)
		envPrefix = strings.ReplaceAll(envPrefix, " ", "_")
		envPrefix = strings.ReplaceAll(envPrefix, "-", "")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config ConfigStruct
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
