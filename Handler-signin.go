package main

import (
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {

	// var to store credentials
	var creds Credentials

	// get user email and password from form value from html
	creds.Email = r.FormValue("email")
	creds.Password = r.FormValue("password")

	//connection to db
	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
	if err != nil {
		//panic("failed to connect database")

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		failedConnectDB += "<h1>Failed to connect to the database</h1>"
		failedConnectDB += "<h4>" + homeLink + "</h4>"

		io.WriteString(w, failedConnectDB)

		return
	}

	// a variable to save the user from the DB
	var user User
	// locates first user with same email
	db.Where("email = ?", creds.Email).First(&user)

	// if entered the wrong password then unauthorised
	if !CheckPasswordHash(creds.Password, user.Password) {

		w.WriteHeader(http.StatusUnauthorized)

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		unauthorized += "<h1>Unsuccessfull sign in - wrong credentials - try again</h1>"
		unauthorized += "<h4>" + homeLink + "</h4>"
		unauthorized += "<h4>" + signUpLink + "</h4>"
		unauthorized += "<h4>" + signInLink + "</h4>"
		io.WriteString(w, unauthorized)
		//io.WriteString(w, "Unsuccessfull - wrong credentials - try again")

		return
	}

	// Declare the expiration time of the token - set to 5 mins
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

		serverError := "<!DOCTYPE HTML><title>Internal server error</title>"
		serverError += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		serverError += "<h1>Internal server errorn</h1>"
		serverError += "<h4>" + homeLink + "</h4>"
		serverError += "<h4>" + signUpLink + "</h4>"
		serverError += "<h4>" + signInLink + "</h4>"
		io.WriteString(w, serverError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	// if all is successfull then .... authorised....
	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
	enterPostLink := "<a href='http://localhost:8000/enterpost.html'>Enter a post</a>"
	displayPostLink := "<a href='http://localhost:8000/displayposts.html'>View a post</a>"
	updatePostLink := "<a href='http://localhost:8000/updatepost.html'>Update a post</a>"
	deletePostLink := "<a href='http://localhost:8000/deletepost.html'>Delete a post</a>"

	authorized := "<!DOCTYPE HTML><title>Authorized</title>"
	authorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
	authorized += "<h1>You are logged in as : " + creds.Email + "</h1>"
	authorized += "<h2>Welcome " + user.Name + "</h2>"
	authorized += "<h4>" + homeLink + "</h4>"
	authorized += "<h4>" + enterPostLink + "</h4>"
	authorized += "<h4>" + displayPostLink + "</h4>"
	authorized += "<h4>" + updatePostLink + "</h4>"
	authorized += "<h4>" + deletePostLink + "</h4>"
	io.WriteString(w, authorized)

}
