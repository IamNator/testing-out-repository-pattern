package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

//CONFIG model for reading configurations
type CONFIG struct {
	DBType     string `json:"_" mapstructure:"DB_TYPE"`
	DBUser     string `json:"_" mapstructure:"DB_USER"`
	DBPassWord string `json:"_" mapstructure:"DB_PASSWORD"`
	DBHost     string `json:"_" mapstructure:"DB_HOST"`
	DBName     string `json:"_" mapstructure:"DB_NAME"`

	PORT string `json:"_" mapstructure:"PORT"`

	AppName          string `mapstructure:"APP_NAME"`
	ValidateTokenURL string `mapstructure:"VALIDATE_TOKEN_URL"`
}

//LoadConfig reads configurations from app.env file
func LoadConfig(path string) (*CONFIG, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	var config CONFIG
	//viper.AutomaticEnv()   //reads environmental variables and overrides those in .env
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return &config, nil
}

//Config is initialized pointer to configuration settings
var Config *CONFIG

func init() {
	cfg, er := LoadConfig("./config")
	if er != nil {
		log.Fatalln(errors.Wrap(er, "unable to read app.env file"))
	}

	Config = cfg
}
