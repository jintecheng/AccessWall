package routes

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func SendJsonPost(jsonBytes []byte, url string) ([]byte, error) {
	//req, err := http.NewRequest("POST", "http://localhost:8001/api/user/new", bytes.NewBuffer(jsonBytes))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	/*
		str := string(body)
		if str == "" {
			panic(str)
		}
	*/

	return body, nil
}
