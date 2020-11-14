package routers

import (
	"errors"
	"strings"

	"github.com/juliankgp/twittBack-Go/db"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/juliankgp/twittBack-Go/models"
)

// Email : email used in all endpoints
var Email string

// IDUser : Id user return from the model
var IDUser string

// ProcessToken : Process to extract the token information
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("Golang_Twitter_React")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := db.CheckExistUser(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
