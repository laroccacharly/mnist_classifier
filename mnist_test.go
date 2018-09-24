package test

import (
	"bytes"
	"encoding/json"
	. "github.com/franela/goblin"
	"github.com/tkanos/gonfig"
	"mnist/main/handlers"
	. "mnist/main/models"
	. "mnist/main/utils"
	"net/http"
	"testing"
)

type Configuration struct {
	Url string
}

func GetConfig() Configuration {
	config := Configuration{}
	err := gonfig.GetConf("config.json", &config)
	if err != nil {
		panic(err)
	}
	return config
}


func assert(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func getConfigURL() string {
	var config Configuration = GetConfig()
	return config.Url
}



func TestConfigUrl(t *testing.T) {
	want := "http://192.168.99.100"
	got := getConfigURL()
	assert(t, got, want)
}

func getNumberSamples() int {
	samples := new([]Sample)
	GetJson(getConfigURL()+"/samples", samples)
	return len(*samples)
}

func classifySample() int {
	payload := handlers.ImagePayload{
		EncodedImage: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABwAAAAcCAAAAABXZoBIAAAAwUlEQVR4nGNgGJqAJe7BiXVRbNgli/89e/bp3/UOKXZMOZtfuxkYpGtv/ft3LRVDeu5HGxDF47r54785aHLavxfBmJr/LqFJGr71gDFj/+9Gk8xYBmOp/vlnhCZ52QLKYD35rw3dPf4cUMaaf4cFMP0CAdG/f6vhkuO+8y8PlxzDxH/XeHDJdf/5p49DirHw398gRhyS3v/+F+Ayk/fgv+lcuOQO/NuPS45h4b9/abjkVL79++WJS1Lr3jpbXHKkAwC+00NH2F/9VQAAAABJRU5ErkJggg==",
	}
	payloadBytes, _ := json.Marshal(payload)
	reader := bytes.NewReader(payloadBytes)
	res, err := http.Post(getConfigURL()+"/classifySample", "application/json", reader)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	label := handlers.LabelPayload{}
	json.NewDecoder(res.Body).Decode(&label)
	return label.Label
}

func TestMnistApi(t *testing.T) {
	g := Goblin(t)
	g.Describe("#samples", func() {
		g.It("should return a list of 10 samples", func() {
			g.Assert(getNumberSamples()).Equal(10)
		})
	})

	g.Describe("#classifySample", func() {
		g.It("should return sample label", func() {
			g.Assert(classifySample()).Equal(7)
		})
	})
}
