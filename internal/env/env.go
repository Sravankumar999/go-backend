package env

import (
	"fmt"
	"os"
)

func GetString(key, fallback string) string {
	fmt.Print(os.Getenv(key))
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val

}
