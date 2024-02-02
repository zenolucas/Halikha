package main

import (
	"Halikha/models"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

	app := fiber.New(fiber.Config{
		BodyLimit: 12 * 1024 * 1024,
	})

	// Serve Bootstrap CSS file
	app.Static("/bootstrap/css", "./node_modules/bootstrap/dist/css")

	// Serve Bootstrap JS files
	app.Static("/bootstrap/js", "./node_modules/bootstrap/dist/js")
	app.Static("/bootstrap/js/popper", "./node_modules/@popperjs/core/dist/umd")

	// Serve other static files
	app.Static("/", "./public")
	app.Static("/", "./public/css")

	app.Get("/", login)

	app.Post("/register", register)

	app.Post("/authenticate", authenticate)

	app.Post("/upload", uploadImage)

	app.Get("/design", design)

	app.Post("/save-artwork", saveArtwork)

	log.Fatal(app.Listen(":3002"))
}

func login(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
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

func uploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return err
	}

	// Save the file to the server
	savePath := "public/uploads/" + file.Filename
	if err := c.SaveFile(file, savePath); err != nil {
		return err
	}

	htmlResponse :=
		`<body>
		<canvas id="canvas"></canvas>

		<script>
		const canvas = new fabric.Canvas('canvas', {
			width: 588,
			height: 647,
		  });
		  
		  // setting canvas bg image
		  canvas.setBackgroundImage('./image-assets/Tshirt-figure.png', canvas.renderAll.bind(canvas));
		  
		  
		  var imgUrl = 'uploads/` + file.Filename + `';
		  fabric.Image.fromURL(imgUrl, function (imgInstance) {
			imgInstance.set({ left: 185, top: 150 });
			canvas.add(imgInstance);
		  });
		  
		  // Define the boundaries for object movement
		  var boundary = {
			left: 180,
			top: 100,
			width: 200,
			height: 325,
		  };
		  
		  // Listen for object moving event
		  canvas.on('object:moving', function (options) {
			var target = options.target;
		  
			// Check if the object is going beyond the boundaries
			if (target.left < boundary.left) {
				target.left = boundary.left;
			}
			if (target.top < boundary.top) {
				target.top = boundary.top;
			}
			if (target.left + target.width > boundary.left + boundary.width) {
				target.left = boundary.left + boundary.width - target.width;
			}
			if (target.top + target.height > boundary.top + boundary.height) {
				target.top = boundary.top + boundary.height - target.height;
			}
		  });
		</script>
	</body>
	`

	return c.SendString(htmlResponse)
}

func design(c *fiber.Ctx) error {
	return c.Redirect("design.html")
}

func saveArtwork(c *fiber.Ctx) error {
	// Parse JSON request body
	var requestBody map[string]string
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	// Get SVG string from the request body
	svgString, ok := requestBody["svgString"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'svgString' in request"})
	}

	// Save SVG data to a file
	filePath := filepath.Join("public/artworks/", "user-drawing.svg")
	err := ioutil.WriteFile(filePath, []byte(svgString), 0644)
	if err != nil {
		log.Println("Error saving file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error saving file"})
	}

	log.Println("File saved successfully")
	return c.JSON(fiber.Map{"success": true, "message": "File saved successfully"})
}
