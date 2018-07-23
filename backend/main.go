package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting backend...")

	data := os.Getenv("BACKEND_DATA")
	if data == "" {
		data = "<empty>"
	}

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, data)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
