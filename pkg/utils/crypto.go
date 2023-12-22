package utils

import "golang.org/x/crypto/bcrypt"

func Crypt(password string) (string, error) {
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
