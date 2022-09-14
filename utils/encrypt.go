package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncryptPassword(password string) string {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		log.Fatal(err)
	}
	return string(encrypt)
}
