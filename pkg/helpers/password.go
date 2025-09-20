package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

/** HashPassword hashes a plain password using bcrypt */
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

/** CheckPassword compares a hashed password with a plain password */
func CheckPassword(hashPassword string, plainPassword []byte) (bool, error) {
	hashPW := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hashPW, plainPassword)
	if err != nil {
		return false, err
	}

	return true, nil
}