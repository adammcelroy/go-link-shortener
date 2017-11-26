package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShorten(t *testing.T) {
	input := "http://google.com"
	expectedGoogle := "https://goo.gl/mR2d"
	expectedBitly := "http://bit.ly/2hF7Z4N"
	expectedTinycc := "http://tiny.cc/vjz4oy"
	fakeParam := "/?a=0"
	channel := make(chan string)

	serverGoogle := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{
			"kind": "urlshortener#url",
			"longUrl": "%v",
			"id": "%v"
		}`, input, expectedGoogle)))
	}))
	serverBitly := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{
			"status_code": 200,
			"status_txt": "OK",
			"data": {
				"long_url": "%v",
				"url": "%v",
				"hash": "2hF7Z4N",
				"global_hash": "3j4ir4",
				"new_hash": 0
			}
		}`, input, expectedBitly)))
	}))
	serverTinycc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{
			"errorCode": "0",
			"errorMessage": "",
			"statusCode":"OK",
			"results": {
				"short_url": "%v",
				"userHash": "vjz4oy",
				"hash": "vjz4oy"
			}
		}`, expectedTinycc)))
	}))

	defer func() {
		serverGoogle.Close()
		serverBitly.Close()
		serverTinycc.Close()
	}()

	API_URL_GOOGLE = serverGoogle.URL + fakeParam
	API_URL_BITLY = serverBitly.URL + fakeParam
	API_URL_TINYCC = serverTinycc.URL + fakeParam

	go shorten(input, API_URL_GOOGLE, channel)
	go shorten(input, API_URL_BITLY, channel)
	go shorten(input, API_URL_TINYCC, channel)

	for link := range channel {
		if link != expectedGoogle && link != expectedBitly && link != expectedTinycc {
			printExpectation(t, "Expected a shortened URL to be returned through the channel")
			printDiscrepancy(t, ("One of: " + expectedGoogle + ", " + expectedBitly + ", " + expectedTinycc), link)
		}
		break
	}
}
