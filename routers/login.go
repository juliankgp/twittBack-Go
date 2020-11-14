package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/jwt"
	"github.com/juliankgp/twittBack-Go/models"
)

// Login : Login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or Pass invalid"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	document, exist := db.TryLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "User or Pass invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Token generation error"+err.Error(), 400)
		return
	}

	resp := models.RespLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: jwtKey, Expires: expirationTime})
}
