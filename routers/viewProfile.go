package routers

import (
	"encoding/json"
	"net/http"

	"github.com/juliankgp/twittBack-Go/db"
)

// ViewProfile : Method to view profile information
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.GetProfile(ID)
	if err != nil {
		http.Error(w, "Error sending the record"+err.Error(), 400)
		return
	}

	w.Header().Set("Contex-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
