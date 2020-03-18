package main

import (
	"encoding/json"
	"net/http"
)

var currentver = ""
var sha = ""
var description = "anz-test-app1"

type Version struct {
	Version     string
	SHA         string
	Description string
}

func main() {
	http.HandleFunc("/version", versionHandle)
	http.ListenAndServe(":8080", nil)
}

func versionHandle(w http.ResponseWriter, r *http.Request) {

	appdetails := Version{currentver, sha, description}
	js, err := json.Marshal(appdetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
