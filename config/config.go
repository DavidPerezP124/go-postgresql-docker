package config

import (
	"os"

	"github.com/joho/godotenv"
)

//Config returns map of env variable required for setup
func Config() (m map[string]string) {
	envVars := make(map[string]string)
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	envVars["DB_USER"] = os.Getenv("DB_USER")
	envVars["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVars["DB_NAME"] = os.Getenv("DB_NAME")
	envVars["PORT"] = os.Getenv("PORT")

	return envVars
}
