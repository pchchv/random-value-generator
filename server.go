package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	r, err := json.Marshal("Flight Service. Version 0.0.1")
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(r)
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
