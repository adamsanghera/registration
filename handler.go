package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../../dbModels/user"
	"github.com/adamsanghera/hashing"
)

func addDefaultHeaders(w http.ResponseWriter, req *http.Request) http.ResponseWriter {
	if acrh, ok := req.Header["Access-Control-Request-Headers"]; ok {
		w.Header().Set("Access-Control-Allow-Headers", acrh[0])
	}
	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := req.Header["Access-Control-Allow-Origin"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := req.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", req.Header["Origin"][0])
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Connection", "Close")
	return w
}

// Register adds a new user to redis, following these steps:
// (1) Parses the json object received in the request.
// (2) Tries to make a new user, following the request.
// (3) Returns the result of this operation as an error (empty message if successful).
func Register(w http.ResponseWriter, req *http.Request) {
	// 0 – setup response
	w = addDefaultHeaders(w, req)
	r := newResponse()
	defer json.NewEncoder(w).Encode(r)

	// 1
	form, err := parseRequest(req)
	r.update(false, err)

	// 2
	salt, hash := hashing.WithNewSalt(form.Password)
	fmt.Println("Registering " + form.Username + " with hash " + hash + " and salt " + salt)
	err = user.Create(form.Username, hash+salt)

	// 3
	r.update(err == nil, err)
}
