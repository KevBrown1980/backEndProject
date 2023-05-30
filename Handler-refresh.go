package main

import (
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func refreshHandler(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
			unauthorized += "<h1>Unsuccessfull post - please check you are signed in</h1>"
			unauthorized += "<h4>" + homeLink + "</h4>"
			unauthorized += "<h4>" + signUpLink + "</h4>"
			unauthorized += "<h4>" + signInLink + "</h4>"
			io.WriteString(w, unauthorized)

			return
		}
		w.WriteHeader(http.StatusBadRequest)
		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		badRequest += "<h1>Bad request</h1>"
		badRequest += "<h4>" + homeLink + "</h4>"

		io.WriteString(w, badRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		unauthorized += "<h1>Unsuccessfull post - please check you are signed in</h1>"
		unauthorized += "<h4>" + homeLink + "</h4>"
		unauthorized += "<h4>" + signUpLink + "</h4>"
		unauthorized += "<h4>" + signInLink + "</h4>"
		io.WriteString(w, unauthorized)

		return
	}
	// (END) The code uptil this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
