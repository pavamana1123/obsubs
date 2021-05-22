package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func updateText(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var textRequest TextRequest
	if json.Unmarshal(reqBody, &textRequest); err != nil {
		w.WriteHeader(500)
		return
	}
	if err = updateOBSText(textRequest.ItemName, textRequest.Text); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}
