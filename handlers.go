package main

// this sneeds to be saved in a .env file*********************************
//var jwtKey = []byte("my_secret_key")

//*********************************************************

// we can score out below as we are using a db
// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

// // Create a struct that models the structure of a user, both in the request body, and in the DB
// type Credentials struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // a struct to store the claim details
// type Claims struct {
// 	Email string `json:"email"`
// 	jwt.RegisteredClaims
// }

//********************************************************
//********* Sign in  *************************************

// func signinHandler(w http.ResponseWriter, r *http.Request) {

// 	// var to store credentials
// 	var creds Credentials

// 	// get user email and password from form value from html
// 	creds.Email = r.FormValue("email")
// 	creds.Password = r.FormValue("password")

// 	//connection to db
// 	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
// 	if err != nil {
// 		//panic("failed to connect database")

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
// 		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		failedConnectDB += "<h1>Failed to connect to the database</h1>"
// 		failedConnectDB += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, failedConnectDB)

// 		return
// 	}

// 	// a variable to save the user from the DB
// 	var user User
// 	// locates first user with same email
// 	db.Where("email = ?", creds.Email).First(&user)

// 	// if entered the wrong password then unauthorised
// 	if !CheckPasswordHash(creds.Password, user.Password) {

// 		w.WriteHeader(http.StatusUnauthorized)

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unauthorized += "<h1>Unsuccessfull sign in - wrong credentials - try again</h1>"
// 		unauthorized += "<h4>" + homeLink + "</h4>"
// 		unauthorized += "<h4>" + signUpLink + "</h4>"
// 		unauthorized += "<h4>" + signInLink + "</h4>"
// 		io.WriteString(w, unauthorized)
// 		//io.WriteString(w, "Unsuccessfull - wrong credentials - try again")

// 		return
// 	}

// 	// Declare the expiration time of the token - set to 5 mins
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	// Create the JWT claims, which includes the username and expiry time
// 	claims := &Claims{
// 		Email: creds.Email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			// In JWT, the expiry time is expressed as unix milliseconds
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	// Declare the token with the algorithm used for signing, and the claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// Create the JWT string
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		// If there is an error in creating the JWT return an internal server error
// 		w.WriteHeader(http.StatusInternalServerError)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 		serverError := "<!DOCTYPE HTML><title>Internal server error</title>"
// 		serverError += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		serverError += "<h1>Internal server errorn</h1>"
// 		serverError += "<h4>" + homeLink + "</h4>"
// 		serverError += "<h4>" + signUpLink + "</h4>"
// 		serverError += "<h4>" + signInLink + "</h4>"
// 		io.WriteString(w, serverError)
// 		return
// 	}

// 	// Finally, we set the client cookie for "token" as the JWT we just generated
// 	// we also set an expiry time which is the same as the token itself
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})

// 	// if all is successfull then .... authorised....
// 	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 	enterPostLink := "<a href='http://localhost:8000/enterpost.html'>Enter a post</a>"
// 	displayPostLink := "<a href='http://localhost:8000/displayposts.html'>View a post</a>"
// 	updatePostLink := "<a href='http://localhost:8000/updatepost.html'>Update a post</a>"
// 	deletePostLink := "<a href='http://localhost:8000/deletepost.html'>Delete a post</a>"

// 	authorized := "<!DOCTYPE HTML><title>Authorized</title>"
// 	authorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 	authorized += "<h1>You are logged in as : " + creds.Email + "</h1>"
// 	authorized += "<h2>Welcome " + user.Name + "</h2>"
// 	authorized += "<h4>" + homeLink + "</h4>"
// 	authorized += "<h4>" + enterPostLink + "</h4>"
// 	authorized += "<h4>" + displayPostLink + "</h4>"
// 	authorized += "<h4>" + updatePostLink + "</h4>"
// 	authorized += "<h4>" + deletePostLink + "</h4>"
// 	io.WriteString(w, authorized)

// }

//********************************************************
//********* Register *************************************

// func signupHandler(w http.ResponseWriter, r *http.Request) {
// 	// var to store credentials
// 	var creds Credentials
// 	// get user email and password from form value from html
// 	creds.Name = r.FormValue("name")
// 	creds.Email = r.FormValue("email")
// 	creds.Password = r.FormValue("password")
// 	// just checking to see if we have them..
// 	fmt.Println(creds)

// 	//connection to db
// 	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
// 	if err != nil {
// 		// io.WriteString(w, "Unsuccesfull - failed to connect to the database")
// 		// panic("failed to connect database")

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
// 		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		failedConnectDB += "<h1>Failed to connect to the database</h1>"
// 		failedConnectDB += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, failedConnectDB)

// 	}

// 	// a variable to save the user from the DB
// 	var user User

// 	// If user not found, create a new user with give conditions
// 	result := db.FirstOrCreate(&user, User{Email: creds.Email, Name: creds.Name})

// 	// hash the password
// 	hashedPassword, _ := HashPassword(creds.Password)

// 	//update user password and email to the DB
// 	user.Password = hashedPassword

// 	// save it to DB
// 	db.Save(&user)

// 	// lets me know if db has been updated - i.e if is been succefull
// 	fmt.Println(result.RowsAffected) // => 1
// 	if result.RowsAffected == 0 {
// 		fmt.Println("That email addresss has already been register")
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"

// 		notSignedUp := "<!DOCTYPE HTML><title>Failed to sign up</title>"
// 		notSignedUp += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		notSignedUp += "<h1>Failed to signed up - The email address may have been used already  - Please try again</h1>"
// 		notSignedUp += "<h4>" + homeLink + "</h4>"
// 		notSignedUp += "<h4>" + signUpLink + "</h4>"

// 		io.WriteString(w, notSignedUp)
// 	} else {
// 		fmt.Println("You have successfully registered")
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"
// 		signedUp := "<!DOCTYPE HTML><title>Successful sign up</title>"
// 		signedUp += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		signedUp += "<h1>Succesfull sign up - now please sign in</h1>"
// 		signedUp += "<h4>" + homeLink + "</h4>"
// 		signedUp += "<h4>" + signInLink + "</h4>"

// 		io.WriteString(w, signedUp)
// 	}

// 	// assignes user passowrd to expected password
// 	//expectedPassword := user.Password

// }

//********************************************************
//********* enterpost *************************************

// func enterPostHandler(w http.ResponseWriter, r *http.Request) {
// 	// We can obtain the session token from the requests cookies, which come with every request
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)
// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign up</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull - are you signed in??? - try again</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)
// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		//signInLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"
//
// 		io.WriteString(w, badRequest)
// 		return
// 	}

// 	// Get the JWT string from the cookie
// 	tknStr := c.Value

// 	// Initialize a new instance of `Claims`
// 	claims := &Claims{}

// 	// Parse the JWT string and store the result in `claims`.
// 	// Note that we are passing the key in this method as well. This method will return an error
// 	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 	// or if the signature does not match
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)

// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign up</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull - are you signed in??? - try again</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"

// 			io.WriteString(w, unauthorized)
// 			return

// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		//signInLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, badRequest)
// 		return
// 	}
// 	if !tkn.Valid {

// 		w.WriteHeader(http.StatusUnauthorized)

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign up</a>"

// 		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unauthorized += "<h1>Unsuccessfull - are you signed in??? - try again</h1>"
// 		unauthorized += "<h4>" + homeLink + "</h4>"
// 		unauthorized += "<h4>" + signUpLink + "</h4>"
// 		unauthorized += "<h4>" + signInLink + "</h4>"
// 		io.WriteString(w, unauthorized)

// 		return
// 	}

// 	// logic for updating posts goes here

// 	// NEED TO PUT |THESE IF STATEMENTS ELSE WHERE
// 	// if not a POST - then do do it....
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/,", http.StatusSeeOther)
// 		return
// 	}

// 	//connection to db
// 	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
// 	if err != nil {
// 		//panic("failed to connect database")
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
// 		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		failedConnectDB += "<h1>Failed to connect to the database</h1>"
// 		failedConnectDB += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, failedConnectDB)
// 		return
// 	}

// 	// temp post to hold the post entry
// 	var post Post
// 	// temp user to hold the User for the entry
// 	var user User

// 	// get post info
// 	// get postentry from html
// 	postentry := r.FormValue("postentry")
// 	// get userID from where claims.Email = user.Email
// 	db.Where("email = ?", claims.Email).Find(&user)
// 	// add values into temp post
// 	post.Postentry = postentry
// 	post.UserID = int(user.ID)

// 	// save the post to DB
// 	db.Save(&post)

// 	//***************************************8s

// 	// Finally, return the welcome message to the user, along with their
// 	// username given in the token
// 	//w.Write([]byte(fmt.Sprintf("Congrats %s! Your posts have been updated", claims.Email)))

// 	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 	enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
// 	displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
// 	successfullPost := "<!DOCTYPE HTML><title>Successful post</title>"
// 	successfullPost += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 	successfullPost += "<h1>Succesfull post</h1>"
// 	successfullPost += "<h4>" + homeLink + "</h4>"
// 	successfullPost += "<h4>" + enterLink + "</h4>"
// 	successfullPost += "<h4>" + displayLink + "</h4>"

// 	io.WriteString(w, successfullPost)

// }

//********************************************************
// //********* Display posts *************************************

// func postsHandler(w http.ResponseWriter, r *http.Request) {
// 	// We can obtain the session token from the requests cookies, which come with every request
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)

// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull post - please check you are signed in</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)

// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, badRequest)
// 		return
// 	}

// 	// Get the JWT string from the cookie
// 	tknStr := c.Value

// 	// Initialize a new instance of `Claims`
// 	claims := &Claims{}

// 	// Parse the JWT string and store the result in `claims`.
// 	// Note that we are passing the key in this method as well. This method will return an error
// 	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 	// or if the signature does not match
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)

// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull post - please check you are signed in</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)

// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, badRequest)
// 		return

// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unauthorized += "<h1>Unsuccessfull post - please check you are signed in</h1>"
// 		unauthorized += "<h4>" + homeLink + "</h4>"
// 		unauthorized += "<h4>" + signUpLink + "</h4>"
// 		unauthorized += "<h4>" + signInLink + "</h4>"
// 		io.WriteString(w, unauthorized)

// 		return

// 	}

// 	// logic for showing posts goes here

// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/,", http.StatusSeeOther)
// 		return
// 	}

// 	//connection to db
// 	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
// 	if err != nil {
// 		//panic("failed to connect database")
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
// 		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		failedConnectDB += "<h1>Failed to connect to the database</h1>"
// 		failedConnectDB += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, failedConnectDB)

// 	}

// 	// temp slics of posts to hold the post entry
// 	var posts []Post
// 	// temp user to hold the User for the entry
// 	var user User

// 	db.Where("email = ?", claims.Email).Find(&user)
// 	// add values into temp post

// 	matchingIDs := int(user.ID)

// 	db.Where("user_id = ?", matchingIDs).Find(&posts)

// 	// a string to capture the result of
// 	resultString := ""

// 	for _, item := range posts {

// 		resultString = resultString + "<p>Post ID: " + strconv.FormatUint(uint64(item.ID), 10) + "</p>" + "<p>" + item.UpdatedAt.String() + "</p>" + "<h2>" + item.Postentry + "</h2>"

// 	}

// 	// // may need to change this to length????
// 	// if resultString == "" {
// 	// 	resultString = "The are no entries found for you"
// 	// }

// 	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 	enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
// 	displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
// 	displayPosts := "<!DOCTYPE HTML><title>Successful post</title>"
// 	displayPosts += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 	displayPosts += "<h1>Posts</h1>"
// 	displayPosts += "<h4>" + resultString + "</h4>"

// 	displayPosts += "<h4>" + homeLink + "</h4>"
// 	displayPosts += "<h4>" + enterLink + "</h4>"
// 	displayPosts += "<h4>" + displayLink + "</h4>"

// 	io.WriteString(w, displayPosts)
// }

//********************************************************
//********* Refresh *************************************

// func refreshHandler(w http.ResponseWriter, r *http.Request) {
// 	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tknStr := c.Value
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// (END) The code uptil this point is the same as the first part of the `Welcome` route

// 	// We ensure that a new token is not issued until enough time has elapsed
// 	// In this case, a new token will only be issued if the old token is within
// 	// 30 seconds of expiry. Otherwise, return a bad request status
// 	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Now, create a new token for the current use, with a renewed expiration time
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the new token as the users `token` cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})
// }

//********************************************************
//********* Logout *************************************

// func logoutHandler(w http.ResponseWriter, r *http.Request) {
// 	// immediately clear the token cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Expires: time.Now(),
// 	})
// 	//time := time.Now().Format("2006-01-02 15:04:05")
// 	// io.WriteString(w, "You are logged out at: \n")
// 	// io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))

// 	// fmt.Println("you are logged out")

// 	homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 	logout := "<!DOCTYPE HTML><title>Logged out</title>"
// 	logout += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 	logout += "<h1>You have logged out</h1>"
// 	logout += "<h4>" + homeLink + "</h4>"
// 	io.WriteString(w, logout)
// }

//************************************************
//**************** delete a post *****************

// func deleteHandler(w http.ResponseWriter, r *http.Request) {

// 	// We can obtain the session token from the requests cookies, which come with every request
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)
// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull delete - please check you are signed in</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)

// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		io.WriteString(w, "Bad request - try again!")
// 		return
// 	}

// 	// Get the JWT string from the cookie
// 	tknStr := c.Value

// 	// Initialize a new instance of `Claims`
// 	claims := &Claims{}

// 	// Parse the JWT string and store the result in `claims`.
// 	// Note that we are passing the key in this method as well. This method will return an error
// 	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 	// or if the signature does not match
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull delete - please check you are signed in</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)

// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, badRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unauthorized += "<h1>Unsuccessfull delete - please check you are signed in</h1>"
// 		unauthorized += "<h4>" + homeLink + "</h4>"
// 		unauthorized += "<h4>" + signUpLink + "</h4>"
// 		unauthorized += "<h4>" + signInLink + "</h4>"
// 		io.WriteString(w, unauthorized)

// 		return

// 	}

// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/,", http.StatusSeeOther)
// 		return
// 	}

// 	//connection to db
// 	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
// 	if err != nil {
// 		//panic("failed to connect database")

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
// 		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		failedConnectDB += "<h1>Failed to connect to the database</h1>"
// 		failedConnectDB += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, failedConnectDB)
// 	}

// 	// temp slics of posts to hold the post entry
// 	var posts []Post
// 	// temp user to hold the User for the entry
// 	var user User

// 	// get  info

// 	db.Where("email = ?", claims.Email).Find(&user)
// 	// add values into temp post

// 	matchingIDs := int(user.ID)

// 	db.Where("user_id = ?", matchingIDs).Find(&posts)

// 	// need to get the deleted id from the html form

// 	postID := r.FormValue("postID")

// 	canDelete := false

// 	for i, v := range posts {
// 		fmt.Println(i, v.ID)
// 		// change v.ID to a string
// 		stringvID := strconv.FormatUint(uint64(v.ID), 10)
// 		if stringvID == postID {
// 			canDelete = true
// 		}

// 		if canDelete {
// 			break
// 		}
// 	}
// 	fmt.Println("Done...")

// 	if canDelete {

// 		db.Where("email = ?", claims.Email).Find(&user)

// 		//delete the expense
// 		db.Delete(&Post{}, postID)
// 		// save cahnges
// 		db.Save(&postID)

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
// 		displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
// 		deleteLink := "<a href='http://localhost:8000/deletepost.html'>Delete posts</a>"
// 		successfullDelete := "<!DOCTYPE HTML><title>Successfull delete</title>"
// 		successfullDelete += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		successfullDelete += "<h1>Delted</h1>"
// 		successfullDelete += "<h4>" + homeLink + "</h4>"
// 		successfullDelete += "<h4>" + enterLink + "</h4>"
// 		successfullDelete += "<h4>" + displayLink + "</h4>"
// 		successfullDelete += "<h4>" + deleteLink + "</h4>"
// 		io.WriteString(w, successfullDelete)
// 	} else {

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
// 		displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
// 		deleteLink := "<a href='http://localhost:8000/deletepost.html'>Delete posts</a>"
// 		unsuccessfullDelete := "<!DOCTYPE HTML><title>Unsuccessfull delete</title>"
// 		unsuccessfullDelete += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unsuccessfullDelete += "<h1>Not deleted</h1>"
// 		unsuccessfullDelete += "<h2>Please try again - check details</h2>"
// 		unsuccessfullDelete += "<h4>" + homeLink + "</h4>"
// 		unsuccessfullDelete += "<h4>" + enterLink + "</h4>"
// 		unsuccessfullDelete += "<h4>" + displayLink + "</h4>"
// 		unsuccessfullDelete += "<h4>" + deleteLink + "</h4>"
// 		io.WriteString(w, unsuccessfullDelete)
// 	}

// }

// func updateHandler(w http.ResponseWriter, r *http.Request) {

// 	// We can obtain the session token from the requests cookies, which come with every request
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)
// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull update - please check you are signed in</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)

// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, badRequest)
// 		return
// 	}

// 	// Get the JWT string from the cookie
// 	tknStr := c.Value

// 	// Initialize a new instance of `Claims`
// 	claims := &Claims{}

// 	// Parse the JWT string and store the result in `claims`.
// 	// Note that we are passing the key in this method as well. This method will return an error
// 	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 	// or if the signature does not match
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 			signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 			signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 			unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 			unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 			unauthorized += "<h1>Unsuccessfull update - please check you are signed in</h1>"
// 			unauthorized += "<h4>" + homeLink + "</h4>"
// 			unauthorized += "<h4>" + signUpLink + "</h4>"
// 			unauthorized += "<h4>" + signInLink + "</h4>"
// 			io.WriteString(w, unauthorized)

// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		badRequest := "<!DOCTYPE HTML><title>BadRequest</title>"
// 		badRequest += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		badRequest += "<h1>Bad request</h1>"
// 		badRequest += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, badRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		signUpLink := "<a href='http://localhost:8000/signup.html'>Sign up</a>"
// 		signInLink := "<a href='http://localhost:8000/signin.html'>Sign in</a>"

// 		unauthorized := "<!DOCTYPE HTML><title>Unauthorized</title>"
// 		unauthorized += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unauthorized += "<h1>Unsuccessfull update - please check you are signed in</h1>"
// 		unauthorized += "<h4>" + homeLink + "</h4>"
// 		unauthorized += "<h4>" + signUpLink + "</h4>"
// 		unauthorized += "<h4>" + signInLink + "</h4>"
// 		io.WriteString(w, unauthorized)

// 		return
// 	}

// 	// if not a POST - then do do it....
// 	if r.Method != "POST" {
// 		http.Redirect(w, r, "/,", http.StatusSeeOther)
// 		return
// 	}

// 	//connection to db
// 	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
// 	if err != nil {
// 		//panic("failed to connect database")

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"

// 		failedConnectDB := "<!DOCTYPE HTML><title>Failed connection</title>"
// 		failedConnectDB += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		failedConnectDB += "<h1>Failed to connect to the database</h1>"
// 		failedConnectDB += "<h4>" + homeLink + "</h4>"

// 		io.WriteString(w, failedConnectDB)
// 	}

// 	// temp slics of posts to hold the post entry
// 	var posts []Post
// 	// temp user to hold the User for the entry
// 	var user User

// 	// get  info

// 	db.Where("email = ?", claims.Email).Find(&user)
// 	// add values into temp post

// 	matchingIDs := int(user.ID)

// 	db.Where("user_id = ?", matchingIDs).Find(&posts)

// 	// need to get the deleted id from the html form

// 	postID := r.FormValue("postID")

// 	// get newupdated post from html
// 	unpdatedEntry := r.FormValue("updatedentry")

// 	// set can update to false before checking if post id belongs to user
// 	canUpdate := false

// 	//************************
// 	//NEED LOGIC FOR CAN update

// 	for i, v := range posts {
// 		fmt.Println(i, v.ID)
// 		// change v.ID to a string
// 		stringvID := strconv.FormatUint(uint64(v.ID), 10)
// 		if stringvID == postID {
// 			canUpdate = true
// 		}

// 		//v = v.ID
// 		if canUpdate {
// 			break
// 		}
// 	}
// 	fmt.Println("Done...")

// 	if canUpdate {

// 		var updatedPost Post

// 		// make sure postID is a unit
// 		uintPostID, _ := strconv.ParseUint(postID, 10, 32)

// 		// get userID from where claims.Email = user.Email
// 		db.Where("id = ?", uintPostID).First(&updatedPost)

// 		//update post
// 		updatedPost.Postentry = unpdatedEntry
// 		fmt.Println("postentry : ", unpdatedEntry)
// 		fmt.Println("postID : ", postID)

// 		db.Save(&updatedPost)

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
// 		displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
// 		deleteLink := "<a href='http://localhost:8000/deletepost.html'>Delete posts</a>"
// 		successfullUpdate := "<!DOCTYPE HTML><title>Successfull update</title>"
// 		successfullUpdate += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		successfullUpdate += "<h1>Updated</h1>"
// 		successfullUpdate += "<h4>" + homeLink + "</h4>"
// 		successfullUpdate += "<h4>" + enterLink + "</h4>"
// 		successfullUpdate += "<h4>" + displayLink + "</h4>"
// 		successfullUpdate += "<h4>" + deleteLink + "</h4>"
// 		io.WriteString(w, successfullUpdate)
// 	} else {

// 		homeLink := "<a href='http://localhost:8000/index.html'>Home</a>"
// 		enterLink := "<a href='http://localhost:8000/enterpost.html'>Sign in</a>"
// 		displayLink := "<a href='http://localhost:8000/displayposts.html'>View posts</a>"
// 		deleteLink := "<a href='http://localhost:8000/deletepost.html'>Delete posts</a>"
// 		unsuccessfullUpdate := "<!DOCTYPE HTML><title>Unsuccessfull update</title>"
// 		unsuccessfullUpdate += "<link rel='stylesheet' type='text/css' href ='style.css'>"
// 		unsuccessfullUpdate += "<h1>Not updated</h1>"
// 		unsuccessfullUpdate += "<h2>Please try again - check details</h2>"
// 		unsuccessfullUpdate += "<h4>" + homeLink + "</h4>"
// 		unsuccessfullUpdate += "<h4>" + enterLink + "</h4>"
// 		unsuccessfullUpdate += "<h4>" + displayLink + "</h4>"
// 		unsuccessfullUpdate += "<h4>" + deleteLink + "</h4>"
// 		io.WriteString(w, unsuccessfullUpdate)
// 	}

// }
