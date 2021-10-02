package env

import (
	"fmt"
	"os"
	"strconv"
)

func MustGetString(name string) string {
	v := os.Getenv(name)
	if v == "" {
		fmt.Printf("environment variable %s was not found or it has an empty string\n", name)
		os.Exit(1)
	}
	return v
}

func GetInt(name string, defaultValue ...int) int {
	v, err := strconv.Atoi(os.Getenv(name))
	if err != nil && len(defaultValue) == 0 {
		fmt.Printf("environment variable %s was not found and no default value was provided\n", name)
		os.Exit(1)
	}

	if err != nil && len(defaultValue) > 0 {
		v = defaultValue[0]
	}

	return v
}
