package routers

import (
	"net/http"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/models"
)

// BreakRelationship function to break relationship
func BreakRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relationship
	t.UserID = IDUser
	t.UserReferenceID = ID

	status, err := db.DeleteRelationship(t)
	if err != nil {
		http.Error(w, "An error occurred deleting the relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "the relationship has not been deleted"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
