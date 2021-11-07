package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	Dbhost              string `mapstructure:"DBHOST"`
	Dbname              string `mapstructure:"DBNAME"`
	Dbuser              string `mapstructure:"DBUSER"`
	Dbpass              string `mapstructure:"DBPASS"`
	Dbport              string `mapstructure:"DBPORT"`
	TelegramBotToken    string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramBotUsername string `mapstructure:"TELEGRAM_BOT_USERNAME"`
	TranslateUrl        string `mapstructure:"TRANSLATE_URL"`
}

func GetConfig() (*Config, error) {

	var err error

	once.Do(func() {

		viper.SetDefault("config", "config.yaml")
		viper.AutomaticEnv()

		config = Config{}

		if err = parseYamlConfig(&config, "config"); err != nil {
			log.Fatal(err)
		}
	})

	return &config, err
}

func parseYamlConfig(cfg interface{}, viperConfigName string) error {
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
