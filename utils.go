package main

import (
	"html/template"
	"strings"
)

func encodeURL(link string) string {
	return template.URLQueryEscaper(link)
}

func enforceProtocol(link string) string {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		return "http://" + link
	}
	return link
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
