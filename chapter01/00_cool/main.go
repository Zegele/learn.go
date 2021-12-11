package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("E:/Geek/src/learn.go")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
