package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func MongoURI() string {
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environmental variables")
	}
	return os.Getenv("MONGODBURI")
}