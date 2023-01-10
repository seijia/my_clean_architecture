package config

import (
	"fmt"
	"path"

	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Net      string
		Host     string
		DBName   string
		Params   struct {
			ParseTime string
		}
	}
	Redis struct {
		Host string
		Port string
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig(configPath string) {

	err := godotenv.Load(path.Join(configPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config := &C

	viper.SetConfigName(os.Getenv("CONFIG_NAME"))
	viper.SetConfigType(os.Getenv("CONFIG_TYPE"))
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
