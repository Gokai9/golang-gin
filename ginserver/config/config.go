package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(en string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error ketika loading env")
	}
	return os.Getenv(en)
}
