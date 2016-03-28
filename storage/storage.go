package storage

import (
	"errors"
)

// Reader reads from storage
type Reader interface {
	Read(string) (string, error)
}

// Writer writes key/value pair into storage
type Writer interface {
	Write(string, string) error
}

// InvalidKeyError returned when key not found in the storage
func InvalidKeyError() error {
	return errors.New("unable to find the url based on the given key")
}
