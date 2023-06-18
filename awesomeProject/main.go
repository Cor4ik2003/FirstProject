package main

import (
	"log"
	"net/http"
)

func handlerRequser(w http.ResponseWriter, r *http.Request) {
	respone := []byte("Hello world!")
	_, err := w.Write(respone)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handlerRequser)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
