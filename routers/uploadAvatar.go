package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/juliankgp/twittBack-Go/db"
	"github.com/juliankgp/twittBack-Go/models"
)

// UploadAvatar upload images
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/avatars/" + IDUser + "." + extension

	funct, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading the image"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(funct, file)
	if err != nil {
		http.Error(w, "Error copying the image"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = db.EditRecord(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error saving the avatar in the database !"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Contex-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
