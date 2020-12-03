package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/juliankgp/twittBack-Go/db"
)

// ListUsers get all users by condition
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "You must send the page parameter as an integer greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.GetAllUsers(IDUser, pag, search, typeUser)
	if !status {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
