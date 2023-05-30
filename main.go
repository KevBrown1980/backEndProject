package main

import (
	"log"
	"net/http"
)

var jwtKey = []byte("my_secret_key")

func main() {
	// create the DB
	createDB()

	// handlers
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/enterpost", enterPostHandler)
	http.HandleFunc("/refresh", refreshHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/update", updateHandler)

	// link to the dir so the golang can access the html and css
	dir := http.Dir("C:/Users/kevin/Desktop/jwt6.1/htdocs")
	fileServer := http.FileServer(dir)
	http.Handle("/", fileServer)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
