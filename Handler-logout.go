package main

import (
	"io"
	"net/http"
	"time"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
	signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
	signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

	logout := "<!DOCTYPE HTML><title>Logged out</title>"
	logout += "<link rel='stylesheet' type='text/css' href ='style.css'>"
	logout += "<h1>You have logged out</h1><br>"
	logout += "<h4>" + homeLink + "</h4>"
	logout += "<h4>" + signUpLink + "</h4>"
	logout += "<h4>" + signInLink + "</h4>"
	io.WriteString(w, logout)
}
