package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func generate(w http.ResponseWriter, r *http.Request) {
	valType := r.URL.Query().Get("type")
	vl := r.URL.Query().Get("length")
	valLength, err := strconv.Atoi(vl)
	if err != nil {
		log.Panic(err)
	}
	val := valueGeneration(valType, valLength)
	v, err := json.Marshal(val)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(v)
	if err != nil {
		log.Panic(err)
	}
}

func retrieve(w http.ResponseWriter, _ *http.Request) {
	// TODO: Getting the value from the id that was returned in the generate method
}

func server() {
	log.Println("Server started!")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/retrieve", retrieve)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
