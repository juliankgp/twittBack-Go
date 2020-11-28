package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juliankgp/twittBack-Go/db"
)

// GetAllTweets get all tweets
func GetAllTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send id parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send page parameter with a value greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	result, ok := db.GetTweets(ID, pag)
	if !ok {
		http.Error(w, "Error reading the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
