package main

import "fmt"

const (
	API_URL_GOOGLE = ("https://www.googleapis.com/urlshortener/v1/url?key=" + API_KEY_GOOGLE)
	API_URL_BITLY  = ("https://api-ssl.bitly.com/v3/shorten?login=" + API_USER_BITLY + "&apiKey=" + API_KEY_BITLY)
	API_URL_TINYCC = ("http://tiny.cc/?c=rest_api&m=shorten&version=2.0.3&login=" + API_USER_TINYCC + "&apiKey=" + API_KEY_TINYCC)
)

func main() {
	var link string

	channel := make(chan string)

	providers := []string{
		API_URL_GOOGLE,
		API_URL_BITLY,
		API_URL_TINYCC,
	}

	fmt.Print("\nEnter a link to shorten: ")
	fmt.Scanln(&link)

	link = enforceProtocol(link)

	for _, provider := range providers {
		go shorten(link, provider, channel)
	}

	fmt.Println("\nShortened:")

	for i, length := 0, len(providers); i < length; i++ {
		fmt.Println(<-channel)
		if i == length-1 {
			fmt.Println("")
		}
	}
}
