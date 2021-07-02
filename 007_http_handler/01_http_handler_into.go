package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("OK\n"))
		if err != nil {
			log.Fatalf("Write failed: %v", err)
		}
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
