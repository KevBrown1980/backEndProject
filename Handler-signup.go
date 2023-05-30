package main

import (
	"fmt"
	"io"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// var to store credentials
	var creds Credentials
	// get user email and password from form value from html
	creds.Name = r.FormValue("name")
	creds.Email = r.FormValue("email")
	creds.Password = r.FormValue("password")
	// just checking to see if we have them..
	fmt.Println(creds)

	//connection to db
	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
	if err != nil {

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		failedConnectDB += "<h1>Failed to connect to the database</h1>"
		failedConnectDB += "<h4>" + homeLink + "</h4>"

		io.WriteString(w, failedConnectDB)

	}

	// a variable to save the user from the DB
	var user User

	// If user not found, create a new user with give conditions
	result := db.FirstOrCreate(&user, User{Email: creds.Email, Name: creds.Name})

	// hash the password
	hashedPassword, _ := HashPassword(creds.Password)

	//update user password and email to the DB
	user.Password = hashedPassword

	// save it to DB
	db.Save(&user)

	// lets me know if db has been updated - i.e if is been succefull
	fmt.Println(result.RowsAffected) // => 1
	if result.RowsAffected == 0 {

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"

		notSignedUp := "<!DOCTYPE HTML><title>Failed to sign up</title>"
		notSignedUp += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		notSignedUp += "<h1>Failed to signed up - The email address may have been used already  - Please try again</h1>"
		notSignedUp += "<h4>" + homeLink + "</h4>"
		notSignedUp += "<h4>" + signUpLink + "</h4>"

		io.WriteString(w, notSignedUp)
	} else {

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"
		signedUp := "<!DOCTYPE HTML><title>Successful sign up</title>"
		signedUp += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		signedUp += "<h1>Succesfull sign up - now please sign in</h1>"
		signedUp += "<h4>" + homeLink + "</h4>"
		signedUp += "<h4>" + signInLink + "</h4>"

		io.WriteString(w, signedUp)
	}

}
