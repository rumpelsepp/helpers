package helpers

import "os"

func MkdirIfNotExists(path string, perm os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, perm)
	}
	return nil
}

func CreateIfNotExists(path string, perm os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
