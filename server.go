package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kr/pretty"
)

type Payload struct {
	TS time.Time `json:"ts"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var payload Payload
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Println(err)
		}
		pretty.Println(payload)
	})
	http.ListenAndServe("127.0.0.1:8082", nil)
}
