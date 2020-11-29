package routers

import (
	"net/http"

	"github.com/juliankgp/twittBack-Go/db"
)

// DeleteTwit route to delete tweet
func DeleteTwit(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send ID parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "An error occurred deleting the tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
