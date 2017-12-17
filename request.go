package register

import (
	"encoding/json"
	"net/http"
)

type requestForm struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func parseRequest(req *http.Request) (requestForm, error) {
	var form requestForm
	return form, json.NewDecoder(req.Body).Decode(&form)
}
