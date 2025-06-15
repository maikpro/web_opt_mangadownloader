package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	ServerPort        string `mapstructure:"SERVER_PORT"`
	OriginUrl         string `mapstructure:"ORIGIN_URL"`
	MongoDBHost       string `mapstructure:"MONGODB_HOST"`
	MongoDBPort       string `mapstructure:"MONGODB_PORT"`
	MongoDBName       string `mapstructure:"MONGODB_NAME"`
	MongoDBCollection string `mapstructure:"MONGODB_COLLECTION"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory: ", err)
	}
	log.Printf("Current working directory is: %s", wd)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
