package main

import "fmt"

const (
	API_URL_GOOGLE = ("https://www.googleapis.com/urlshortener/v1/url?key=" + API_KEY_GOOGLE)
	API_URL_BITLY  = ("https://api-ssl.bitly.com/v3/shorten?login=" + API_USER_BITLY + "&apiKey=" + API_KEY_BITLY)
	API_URL_TINYCC = ("http://tiny.cc/?c=rest_api&m=shorten&version=2.0.3&login=" + API_USER_TINYCC + "&apiKey=" + API_KEY_TINYCC)
)

func main() {
	var link string

	// Create channel for communication between subroutines
	channel := make(chan string)

	// Create a slice containing our API URLs
	providers := []string{
		API_URL_GOOGLE,
		API_URL_BITLY,
		API_URL_TINYCC,
	}

	// Get the URL we would like the shorten
	fmt.Print("\nEnter a link to shorten: ")
	fmt.Scanln(&link)

	// Add http:// to beginning of the inputted
	// URL if has no leading protocol already
	link = enforceProtocol(link)

	// Start a subroutine for each provider
	for _, provider := range providers {
		go shorten(link, provider, channel)
	}

	// Show results
	fmt.Println("\nSHORTENED:")

	// Wait to hear back from each subroutine
	// and print the results
	for i, length := 0, len(providers); i < length; i++ {
		fmt.Println(<-channel)

		if i == length-1 {
			fmt.Println("")
		}
	}
}
