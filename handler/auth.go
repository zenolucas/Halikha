package handler

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

			// login success!
			// Get session from storage
			// sess, err := store.Get(c)
			/*
				if err != nil {
					panic(err)
				}
			*/

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