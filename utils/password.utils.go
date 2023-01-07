package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(candidatePassword string, arrivedPassword string) interface{} {
	arrivedHashedPassword, err := bcrypt.GenerateFromPassword([]byte(arrivedPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("could not hash password %w", err)
	}

	if string(arrivedHashedPassword) == candidatePassword {
		return nil
	}
	return err
	//return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
