package tools

import (
	"github.com/subosito/gotenv"
	"os"
	"strconv"
)

// Env reads specified environment variable. If no value has been found,
// fallback is returned.
func Env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func EnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		} else {
			panic(err)
		}
	}
	return fallback
}



// LoadEnvFile loads environment variables defined in an .env formatted file.
func LoadEnvFile(envfilepath string) error {
	err := gotenv.Load(envfilepath)
	return err
}