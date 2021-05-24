package helpers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	if err := godotenv.Load(RootDir() + "/.env"); err != nil {
		fmt.Println(err)
		log.Fatal("error loading .env file")
	}
}
