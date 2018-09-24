package main

import (
	"mnist/main/handlers"
	"net/http"
)

func main() {
	// make client
	client := handlers.ClassifierClient{URL: "http://classifier"}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/samples", client.GetSamples)
	http.HandleFunc("/classifySample", client.ClassifySample)
	http.ListenAndServe(":100", nil)
}
