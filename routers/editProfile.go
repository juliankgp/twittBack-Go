package routers

import (
	"encoding/json"
	"net/http"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/models"
)

// EditProfile function to edit profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data"+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.EditRecord(t, IDUser)
	if err != nil {
		http.Error(w, "Error modifying registry, please try again"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error modifying registry", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
