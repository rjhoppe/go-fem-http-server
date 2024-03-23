package api

import (
	"encoding/json"
	"net/http"

	"github.com/rjhoppe/go-http-server/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// make sure method is POST
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		w.WriteHeader(http.StatusCreated)
		// Could uncomment below as well
		// w.Write([]byte("Success"))
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
