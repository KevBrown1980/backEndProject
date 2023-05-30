package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
			unauthorized += "<h1>Unsuccessfull delete - please check you are signed in</h1>"
			unauthorized += "<h4>" + homeLink + "</h4>"
			unauthorized += "<h4>" + signUpLink + "</h4>"
			unauthorized += "<h4>" + signInLink + "</h4>"
			io.WriteString(w, unauthorized)

			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad request - try again!")
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
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
			unauthorized += "<h1>Unsuccessfull delete - please check you are signed in</h1>"
			unauthorized += "<h4>" + homeLink + "</h4>"
			unauthorized += "<h4>" + signUpLink + "</h4>"
			unauthorized += "<h4>" + signInLink + "</h4>"
			io.WriteString(w, unauthorized)

			return
		}
		w.WriteHeader(http.StatusBadRequest)
		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		badRequest += "<h1>Bad request</h1>"
		badRequest += "<h4>" + homeLink + "</h4>"
		badRequest += "<h4>" + signUpLink + "</h4>"
		badRequest += "<h4>" + signInLink + "</h4>"

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
		unauthorized += "<h1>Unsuccessfull delete - please check you are signed in</h1>"
		unauthorized += "<h4>" + homeLink + "</h4>"
		unauthorized += "<h4>" + signUpLink + "</h4>"
		unauthorized += "<h4>" + signInLink + "</h4>"
		io.WriteString(w, unauthorized)

		return

	}

	if r.Method != "POST" {
		http.Redirect(w, r, "/,", http.StatusSeeOther)
		return
	}

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
	}

	// temp slics of posts to hold the post entry
	var posts []Post
	// temp user to hold the User for the entry
	var user User

	// get  info

	db.Where("email = ?", claims.Email).Find(&user)
	// add values into temp post

	matchingIDs := int(user.ID)

	db.Where("user_id = ?", matchingIDs).Find(&posts)

	// need to get the deleted id from the html form

	postID := r.FormValue("postID")

	canDelete := false

	for i, v := range posts {
		fmt.Println(i, v.ID)
		// change v.ID to a string
		stringvID := strconv.FormatUint(uint64(v.ID), 10)
		if stringvID == postID {
			canDelete = true
		}

		if canDelete {
			break
		}
	}
	fmt.Println("Done...")

	if canDelete {

		db.Where("email = ?", claims.Email).Find(&user)

		//delete the expense
		db.Delete(&Post{}, postID)
		// save cahnges
		db.Save(&postID)

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		enterLink := "<a href='http://localhost:8000/enterpost.html'>Create a post</a>"
		displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
		deleteLink := "<a href='http://localhost:8000/deletepost.html'>Delete posts</a>"
		successfullDelete := "<!DOCTYPE HTML><title>Successfull delete</title>"
		successfullDelete += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		successfullDelete += "<h1>Delted</h1>"
		successfullDelete += "<h4>" + homeLink + "</h4>"
		successfullDelete += "<h4>" + enterLink + "</h4>"
		successfullDelete += "<h4>" + displayLink + "</h4>"
		successfullDelete += "<h4>" + deleteLink + "</h4>"
		io.WriteString(w, successfullDelete)
	} else {

		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
		enterLink := "<a href='http://localhost:8000/enterpost.html'>Create a post</a>"
		displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
		deleteLink := "<a href='http://localhost:8000/deletepost.html'>Delete posts</a>"
		unsuccessfullDelete := "<!DOCTYPE HTML><title>Unsuccessfull delete</title>"
		unsuccessfullDelete += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		unsuccessfullDelete += "<h1>Not deleted</h1>"
		unsuccessfullDelete += "<h2>Please try again - check details</h2>"
		unsuccessfullDelete += "<h4>" + homeLink + "</h4>"
		unsuccessfullDelete += "<h4>" + enterLink + "</h4>"
		unsuccessfullDelete += "<h4>" + displayLink + "</h4>"
		unsuccessfullDelete += "<h4>" + deleteLink + "</h4>"
		io.WriteString(w, unsuccessfullDelete)
	}

}
