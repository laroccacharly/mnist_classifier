package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mnist/main/models"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGetSamples(t *testing.T) {
	// first I receive a request from the browser
	// The Request is sent to a handler
	// the handler makes a Request to the Classifier Server
	// The server responds with a a list of samples in a json
	// the handler returns this list to the client

	// Generate some expected data from server
	sampleList := models.GenerateSampleList()
	sampleListBytes, _ := json.Marshal(&sampleList)

	// setup classifier server, to return a list a samples when called
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sampleListBytes)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()

	// Request made from client
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	recorder := httptest.NewRecorder()


	// client init that will call the server
	client := ClassifierClient{URL: srv.URL}

	// client.GetSamples
	client.GetSamples(recorder, req)

	// Get the response from the recorder
	resp := recorder.Result()
	receiveSampleList := []models.Sample{}
	json.NewDecoder(resp.Body).Decode(&receiveSampleList)

	// assert the json is a list of 10 samples.SampleList
	if len(receiveSampleList) != 10 {
		t.Errorf("Error sample list not the right len")
	}

}


func TestClassifySample(t *testing.T) {
	// Receive a POST request from client
	// Make the body of the request
	postBody := ImagePayload{EncodedImage: "image123"}
	body, _ := json.Marshal(&postBody)
	reader := bytes.NewReader(body)
	// Make the request
	req := httptest.NewRequest("POST", "http://example.com/foo", reader)

	// Make a recoder
	recorder := httptest.NewRecorder()

	// Make the server that will be called by our handler
	// The response from the server will be a label
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request from my client
		payload := ImagePayload{}
		json.NewDecoder(r.Body).Decode(&payload)
		payloadBytes, _ := json.Marshal(payload)
		fmt.Printf("Received on test server image payload : %s", string(payloadBytes))
		if  payload.EncodedImage == "" {
			panic("did not receive a correct encoded image")
		}
		// Return some test data
		label := LabelPayload{Label: 7}
		labelBytes, _ := json.Marshal(&label)
		w.Write(labelBytes)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	// make client
	client := ClassifierClient{URL: srv.URL}

	// Send the the request and recorder to handler
	client.ClassifySample(recorder, req)

	// Get the response on the recorder
	resp := recorder.Result()
	receivedLabel := LabelPayload{}
	json.NewDecoder(resp.Body).Decode(&receivedLabel)

	if receivedLabel.Label != 7 {
		t.Errorf("Oops did not receive the right label")
	}
}

