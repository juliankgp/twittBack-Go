package middlew

import (
	"net/http"

	"github.com/juliankgp/twittBack-Go/db"
)

// CheckDB : Interceptor to verify the connection with the DB
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Lost connection to the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
