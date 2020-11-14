package db

import (
	"github.com/juliankgp/twittBack-Go/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin : Check if user can login
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckExistUser(email)
	if found == false {
		return user, false
	}

	passBytes := []byte(password)

	passDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passDB, passBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
