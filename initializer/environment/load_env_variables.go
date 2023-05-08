package environment

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro to loading .env file")
	}

}
