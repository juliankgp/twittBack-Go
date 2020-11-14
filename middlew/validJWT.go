package middlew

import (
	"github.com/juliankgp/twittBack-Go/routers"
)

// ValidJWT : Valid the JWT
func ValidJWT(next http.HandlerFunc) http.HandlerFunc  {
	return func (w http.ResponseWriter, r *http.Request){
		_,_,_; err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in Token !", + err.Error(), http.StatusBadRequest)
			return
		}
	}
	next.ServeHTPP(w,r)
}