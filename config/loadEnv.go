package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) error {
	fileName := fmt.Sprintf(".env.%s", env)
	return godotenv.Load(fileName)
}
