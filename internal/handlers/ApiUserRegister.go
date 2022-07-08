package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ApiUserRegister(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(b))

	type JSONIn struct {
		Login2Reg string `json:"login"`
		Pass2Reg  string `json:"password"`
	}
	var data JSONIn
	err = json.Unmarshal(b, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(data.Login2Reg)

	return
}
