package routers

import (
	"encoding/json"
	"net/http"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/models"
)

// SignIn : Function to register a new user in the db
func SignIn(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password must be at least 6 characters long", 400)
		return
	}

	_, found, _ := db.CheckExistUser(t.Email)
	if found == true {
		http.Error(w, "The user already exist", 400)
		return
	}

	_, status, err := db.InsertRecord(t)
	if err != nil {
		http.Error(w, "Failed to save user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Status False, failed to save user"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
