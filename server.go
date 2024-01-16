package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"Halikha/models"
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
		DBName:               "test",
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

	// Authentication route
	app.Post("/authenticate", func(c *fiber.Ctx) error {
		// Retrieve form data
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Simple authentication (replace with your actual authentication logic)



		var user models.User
		// to replace the code above for authentication logic
		row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
		if err := row.Scan(&user.ID, &user.Username, &user.Usertype, &user.Password); err != nil {
			return c.SendString("wrong username or password")
		}

		fmt.Println(user.ID, user.Username, user.Password)

		if user.Username == username {
			if user.Password == password {

				if user.Usertype == "artist" {
					return c.SendFile("public/artistHomePage.html")
				} else {
					return c.SendFile("public/customerHomePage.html")
				}
			}
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}

func login(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
}
