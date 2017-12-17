package register

import "errors"
import "fmt"

type response struct {
	Successful bool  `json:"Successful"`
	ErrMsg     error `json:"ErrMsg"`
}

func (r *response) update(s bool, err error) {
	r.Successful = s
	r.ErrMsg = err
	if err != nil {
		fmt.Println(err)
	}
	if err != nil && err.Error() != "That username already exists" {
		panic(err)
	}
}

func newResponse() *response {
	resp := response{
		ErrMsg: errors.New("Unknown error"),
	}
	return &resp
}
