package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juliankgp/twittBack-Go/middlew"
	"github.com/juliankgp/twittBack-Go/routers"
	"github.com/rs/cors"
)

// Controller : func to take the control to the app
func Controller() {
	router := mux.NewRouter()

	router.HandleFunc("/signIn", middlew.CheckDB(routers.SignIn)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewProfile", middlew.CheckDB(middlew.ValidJWT(routers.ViewProfile))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Println("Listening on port: " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
