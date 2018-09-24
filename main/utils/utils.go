package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetRequestBody(r *http.Request) []byte {
	defer r.Body.Close()
	resBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	return resBytes
}

func CheckExists(path string) (bool, error) {
	_, err := os.Stat(path)
	fmt.Println(path)

	if err == nil {
		fmt.Println("exists")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("not exists")
		return false, nil
	}

	fmt.Println("exists")
	return true, err
}

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Got an error calling url : " + url)
		fmt.Println("Error message : " + err.Error())
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
