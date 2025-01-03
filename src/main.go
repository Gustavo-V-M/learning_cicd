package main

import (
	"http_example/page"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", page.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
