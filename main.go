package main

import (
	"fmt"
	"gapi/infrastructure/persistence"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	services, _ := persistence.NewRepositories()
	user, _ := services.User.GetUser("1")
	fmt.Println(user.ID)
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
