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
	TelegramBotDebug    bool   `mapstructure:"TELEGRAM_BOT_DEBUG"`
	TranslateUrl        string `mapstructure:"TRANSLATE_URL"`
}

func GetConfig() (*Config, error) {

	var err error

	once.Do(func() {

		viper.AutomaticEnv()

		config = Config{}

		if err = parseYamlConfig(&config, "config.yaml"); err != nil {
			log.Fatal(err)
		}
	})

	return &config, err
}

func parseYamlConfig(cfg interface{}, fileName string) error {
	if len(fileName) == 0 {
		return errors.New("config file not found")
	}

	filePath, err := filepath.Abs(fileName)

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
