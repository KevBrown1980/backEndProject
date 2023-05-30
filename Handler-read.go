package main

import (
	"io"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func postsHandler(w http.ResponseWriter, r *http.Request) {
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
			unauthorized += "<h1>Unsuccessfull post - please check you are signed in</h1>"
			unauthorized += "<h4>" + homeLink + "</h4>"
			unauthorized += "<h4>" + signUpLink + "</h4>"
			unauthorized += "<h4>" + signInLink + "</h4>"
			io.WriteString(w, unauthorized)

			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
		badRequest += "<h1>Bad request</h1>"
		badRequest += "<h4>" + homeLink + "</h4>"

		io.WriteString(w, badRequest)
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

	// logic for showing posts goes here

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

	db.Where("email = ?", claims.Email).Find(&user)
	// add values into temp post

	matchingIDs := int(user.ID)

	db.Where("user_id = ?", matchingIDs).Find(&posts)

	// a string to capture the result of
	resultString := ""

	for _, item := range posts {

		resultString = resultString + "<p>Post ID: " + strconv.FormatUint(uint64(item.ID), 10) + "</p>" + "<p>" + item.UpdatedAt.String() + "</p>" + "<h2>" + item.Postentry + "</h2>"

	}

	// // may need to change this to length????
	// if resultString == "" {
	// 	resultString = "The are no entries found for you"
	// }

	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
	enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
	displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
	displayPosts := "<!DOCTYPE HTML><title>Successful post</title>"
	displayPosts += "<link rel='stylesheet' type='text/css' href ='style.css'>"
	displayPosts += "<h1>Posts</h1>"
	displayPosts += "<h4>" + resultString + "</h4>"

	displayPosts += "<h4>" + homeLink + "</h4>"
	displayPosts += "<h4>" + enterLink + "</h4>"
	displayPosts += "<h4>" + displayLink + "</h4>"

	io.WriteString(w, displayPosts)
}
