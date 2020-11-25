package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/models"
)

// SaveTweet endpoint to save tweets
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	bodyTweet := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(bodyTweet)
	if err != nil {
		http.Error(w, "An error occurred while inserting the record, please try again"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error inserting the tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
