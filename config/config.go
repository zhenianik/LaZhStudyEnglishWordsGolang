package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	JdbcDriver          string `mapstructure:"JDBC_DRIVER"`
	JdbcUrl             string `mapstructure:"JDBC_URL"`
	JdbcUsername        string `mapstructure:"JDBC_USERNAME"`
	JdbcPassword        string `mapstructure:"JDBC_PASSWORD"`
	TelegramBotToken    string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramBotUsername string `mapstructure:"TELEGRAM_BOT_USERNAME"`
}

func GetConfig() (Config, error) {

	var err error

	viper.SetDefault("config", "config.yaml")
	viper.AutomaticEnv()

	config := Config{}

	if err := ParseYamlConfig(&config, "config"); err != nil {
		return config, err
	}

	return config, err
}

func ParseYamlConfig(cfg interface{}, viperConfigName string) error {
	filePath := viper.GetString(viperConfigName)
	if len(filePath) == 0 {
		return errors.New("config file not found")
	}

	filePath, err := filepath.Abs(filePath)

	if err != nil {
		return err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("config file '%s' not found", filePath)
	}

	v := viper.New()

	v.SetConfigFile(filePath)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	return v.Unmarshal(&cfg)

}
