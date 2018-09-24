package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type ImagePayload struct {
	EncodedImage string `json:"encoded_image"`
}

type LabelPayload struct {
	Label int `json:"label"`
}


type ClassifierClient struct {
	URL string
}

func (c *ClassifierClient) GetSamples(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting samples from classifier")
	resp, err := http.Get(c.URL + "/samples")
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	enableCors(&w)
	w.Write(body)
}


func (c *ClassifierClient) ClassifySample(w http.ResponseWriter, r *http.Request) {
	payloadBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(payloadBytes)
	fmt.Printf("Payload Received in ClassifySample : %s", string(payloadBytes))
	resp, err := http.Post(c.URL + "/classifySample", "application/json", reader)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}


