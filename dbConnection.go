package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createDB() {

	db, err := gorm.Open(sqlite.Open("post.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Check if table exists and drop it if found
	// db.Migrator().DropTable(&User{})
	// db.Migrator().DropTable(&Post{})

	// Create tables
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})

	//create multiple users
	// db.Create(&User{Email: "chris@barclays.com", Password: "1234"})
	// db.Create(&User{Email: "veronica@barclays.com", Password: "1234"})
	// db.Create(&User{Email: "karolina@barclays.com", Password: "1234"})
	// db.Create(&User{Email: "rodrigo@barclays.com", Password: "1234"})
	// db.Create(&User{Email: "kevin@barclays.com", Password: "1234"})

	// Create entries posts
	// db.Create(&Post{Blog: "qqqqqqqqqq", UserID: 1})
	// db.Create(&Post{Blog: "dddddddddddd", UserID: 3})
	// db.Create(&Post{Blog: "wwwwwwwwww", UserID: 2})
	// db.Create(&Post{Blog: "eeeeeeeeeee", UserID: 2})
	// db.Create(&Post{Blog: "rrrrrrrrrr", UserID: 3})
	// db.Create(&Post{Blog: "tttttttttttt", UserID: 4})
	// db.Create(&Post{Blog: "yyyyyyyyyy", UserID: 5})
	// db.Create(&Post{Blog: "uuuuuuuuuuuuu", UserID: 1})
	// db.Create(&Post{Blog: "jjjjjjjjjjjj", UserID: 2})
	// db.Create(&Post{Blog: "nnnnnnnnnnnnn", UserID: 3})
	// db.Create(&Post{Blog: "fffffffffffff", UserID: 4})
	// db.Create(&Post{Blog: "mmmmmmmmmmmmmmmmmmm", UserID: 3})

}
