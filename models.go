package main

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// creates a one user to many posts relationship
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Posts    []Post
}

// posts are the users entries
type Post struct {
	gorm.Model
	Postentry string
	UserID    int
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// a struct to store the claim details
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
