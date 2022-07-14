package main

import (
	"log"
	"net/http"
	"strconv"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(getJSON("", "Random Value Service. Version 0.1"))
	if err != nil {
		log.Panic(err)
	}
}

func generate(w http.ResponseWriter, r *http.Request) {
	valLength := 0
	valType := r.URL.Query().Get("type")
	vl := r.URL.Query().Get("length")
	if vl != "" {
		valLength, _ = strconv.Atoi(vl)
	}
	val := valueGeneration(valType, valLength)
	v := getJSON("value: ", val)
	id := getJSON("id: ", toDB(val))
	v = append(v, id...)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(v)
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
