package utils

import uuid "github.com/satori/go.uuid"

func GenerateID() string {
	return uuid.NewV4().String()
}
