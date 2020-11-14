package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/juliankgp/twittBack-Go/models"
)

// GenerateJWT : Generate the encrypt with JWT
func GenerateJWT(t models.User) (string, error) {

	myKey := []byte("Golang_Twitter_React")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.LastName,
		"birthdate": t.Birthdate,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
