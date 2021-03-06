package db

import "golang.org/x/crypto/bcrypt"

// EncryptPass : Func to encrypt the password
func EncryptPass(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
