package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashed)
}

func ComparePassword(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
