package main

import (
	"Halikha/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

func main() {
	// connect to DB first
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "Halikha",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	app := fiber.New()

	// Serve Bootstrap CSS file
	app.Static("/bootstrap/css", "./node_modules/bootstrap/dist/css")

	// Serve Bootstrap JS files
	app.Static("/bootstrap/js", "./node_modules/bootstrap/dist/js")
	app.Static("/bootstrap/js/popper", "./node_modules/@popperjs/core/dist/umd")

	// Serve other static files
	app.Static("/", "./public")

	app.Get("/", login)

	app.Post("/register", register)

	// Authentication route
	app.Post("/authenticate", authenticate)

	log.Fatal(app.Listen(":3001"))
}

func login(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
}

func test(c *fiber.Ctx) error {
	return c.SendFile("public/test.html")
}

func authenticate(c *fiber.Ctx) error {
	// Retrieve form data
	username := c.FormValue("username")
	password := c.FormValue("password")
	var user models.User

	row := db.QueryRow("SELECT * FROM Users WHERE Username = ?", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Usertype); err != nil {
		return c.SendString("error in retrieving data from database")
	}

	if user.Username == username {
		if user.Password == password {

			if user.Usertype == "Artist" {
				return c.Redirect("artistHomePage.html")
			} else {
				return c.Redirect("customerHomePage.html")
			}
		}
	}
	return c.Redirect("/")
}

func register(c *fiber.Ctx) error {
	// Retrieve form data
	email := c.FormValue("email")
	username := c.FormValue("username")
	usertype := c.FormValue("usertype")
	password := c.FormValue("password")

	// 	to be implemented: hashing the password

	// Insert user into the database
	_, err := db.Exec("INSERT INTO Users (Username, PasswordHash, Email, UserType ) VALUES (?, ?, ?, ?)", username, password, email, usertype)
	if err != nil {
		fmt.Errorf("error in registering user: %v", err)
		return err
	}

	// Redirect to the login page
	return c.Redirect("/")
}
