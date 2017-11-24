package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func shorten(link string, provider string, channel chan string) {
	var requestData map[string]string
	var requestDataForPost []byte
	var requestDataForGet string
	var requestType string
	var req *http.Request
	var err error

	if provider == API_URL_GOOGLE {
		requestType = "POST"
		requestData = map[string]string{
			"longUrl": link,
		}
	} else if provider == API_URL_BITLY {
		requestType = "GET"
		requestData = map[string]string{
			"longUrl": encodeURL(link),
			"domain":  "bit.ly",
		}
	} else if provider == API_URL_TINYCC {
		requestType = "GET"
		requestData = map[string]string{
			"longUrl": encodeURL(link),
		}
	}

	if requestType == "POST" {
		requestDataForPost, _ = json.Marshal(requestData)
		req, err = http.NewRequest(requestType, provider, bytes.NewBuffer(requestDataForPost))
		req.Header.Set("Content-Type", "application/json")

	} else if requestType == "GET" {
		for key, value := range requestData {
			requestDataForGet += ("&" + key + "=" + value)
		}
		req, err = http.NewRequest(requestType, (provider + requestDataForGet), bytes.NewBuffer([]byte{}))
	}

	client := &http.Client{}
	res, err := client.Do(req)
	panicIfErrors(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	panicIfErrors(err)

	if provider == API_URL_GOOGLE {
		var google GoogleShortenedLink
		err = json.Unmarshal(body, &google)
		panicIfErrors(err)
		channel <- google.Value
	} else if provider == API_URL_BITLY {
		var bitly BitlyShortenedLink
		err = json.Unmarshal(body, &bitly)
		panicIfErrors(err)
		channel <- bitly.Data.Value
	} else if provider == API_URL_TINYCC {
		var tinycc TinyccShortenedLink
		err = json.Unmarshal(body, &tinycc)
		panicIfErrors(err)
		channel <- tinycc.Data.Value
	}
}