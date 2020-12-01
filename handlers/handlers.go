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
	router.HandleFunc("/viewprofile", middlew.CheckDB(middlew.ValidJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/editProfile", middlew.CheckDB(middlew.ValidJWT(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweets", middlew.CheckDB(middlew.ValidJWT(routers.GetAllTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidJWT(routers.DeleteTwit))).Methods("DELETE")

	router.HandleFunc("/avatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/avatar", middlew.CheckDB(middlew.ValidJWT(routers.UploadAvatar))).Methods("POST")

	router.HandleFunc("/banner", middlew.CheckDB(routers.GetBanner)).Methods("GET")
	router.HandleFunc("/banner", middlew.CheckDB(middlew.ValidJWT(routers.UploadBanner))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Println("Listening on port: " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
