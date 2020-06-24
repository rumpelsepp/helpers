package helpers

import (
	"os"
	"strconv"
)

func GetEnvBool(name string) bool {
	if rawVal, ok := os.LookupEnv(name); ok {
		if val, err := strconv.ParseBool(rawVal); val && err == nil {
			return val
		}
	}
	return false
}
