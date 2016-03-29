package storage

import (
	"errors"
	"fmt"
)

// Storage is both reader and writer
type Storage interface {
	Reader
	Writer
}

// Reader reads from storage
type Reader interface {
	Read(string) (string, error)
}

// Writer writes the url into storage, and return its short version
type Writer interface {
	Write(string) (string, error)
}

// InvalidKeyError returned when key not found in the storage
func InvalidKeyError(key string) error {
	return errors.New(fmt.Sprintf("unable to find the value for key: %s", key))
}
