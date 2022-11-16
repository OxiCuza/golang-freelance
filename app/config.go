package app

import (
	"github.com/joho/godotenv"
	"golang-freelance/helper"
	"os"
)

func EnvVariable(key string) string {
	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	return os.Getenv(key)
}
