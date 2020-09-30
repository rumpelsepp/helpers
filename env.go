package helpers

import (
	"os"
	"path/filepath"
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

func ConfigPath(name string) string {
	p, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(p, name)
}
