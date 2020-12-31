package utils

import (
	"errors"
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Printf("Error: %s", err)
	}
}

func CreateAndLogError(message string) error {
	err := errors.New(message)
	LogError(err)
	return err
}
