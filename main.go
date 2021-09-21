package main

import (
	"fmt"
	"net/http"
)

type Handler interface {
	GetTitle() string
}

type Page struct{}

func (p Page) GetTitle() string {
	return "Hello"
}

func hello(w http.ResponseWriter, req *http.Request) {
	page := Page{}
	fmt.Println(page.GetTitle())
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
