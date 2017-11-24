package main

// GoogleShortenedLink contains return value from API
type GoogleShortenedLink struct {
	Value string `json:"id"`
}
type BitlyShortenedLink struct {
	Data struct {
		Value string `json:"url"`
	} `json:"data"`
}
type TinyccShortenedLink struct {
	Data struct {
		Value string `json:"short_url"`
	} `json:"results"`
}
