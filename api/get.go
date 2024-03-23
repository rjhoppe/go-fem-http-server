package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rjhoppe/go-http-server/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// api/exhibitions?id=34
	id := r.URL.Query()["id"]
	// need to convert id into int and confirm it is an int
	if id != nil { // will try to serve only one
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode((data.GetAll()))
	}
}
