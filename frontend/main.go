package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const backendURL = "http://backend/data"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(backendURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "the result is %s", b)
}

func main() {
	log.Println("Starting frontend...")
	http.HandleFunc("/view", viewHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
