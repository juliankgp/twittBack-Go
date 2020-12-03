package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juliankgp/twittBack-Go/db"
)

// GetFollowersTweets get all follower tweets
func GetFollowersTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send page parameter", http.StatusBadRequest)
		return
	}

	pagTemp, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send page parameter as integer more than zero", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, okay := db.GetTweetsFollowers(IDUser, pag)
	if !okay {
		http.Error(w, "Error reading the tweets", http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
