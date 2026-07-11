package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() { 
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Falling back to system environment variables.")
	}
}

func init() {
    // Try to load .env early so package-level initializers can read env vars.
    _ = godotenv.Load()
}
func GetInt(key string , fallback int) int{

	value,ok := os.LookupEnv(key)
	if !ok{
		return fallback
	}
	intValue,err  := strconv.Atoi(value)

	if err != nil{
		log.Fatal(err)
	}

	return intValue

}


func GetString(key string , fallback string) string{

	value,ok := os.LookupEnv(key)
	if !ok{
		return fallback
	}


	
	return value

}