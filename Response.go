package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Answer string `json:"answer"`
	Image  string `json:"image"`
}

func GetAnswer() (*Response, error) {
	resp, err := http.Get("https://yesno.wtf/api")
	if err != nil {
		return nil, err
	}

	return readHttpResponse(resp)
}

func GetForcedAnswer(ans string) (*Response, error) {
	resp, err := http.Get("https://yesno.wtf/api?force=" + ans)
	if err != nil {
		return nil, err
	}

	return readHttpResponse(resp)
}

func readHttpResponse(resp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Response
	err = json.Unmarshal(body, &r)
	return &r, err
}
