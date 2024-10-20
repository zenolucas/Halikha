package handler

import "net/http"

func uploadImage(w http.ResponseWriter, r *http.Request) error {
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
		<canvas id="canvas"></canvas>design

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
		  canvas.on('object:moving', fgunction (options) {
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

	// filename will be UserID_timestamp
	// timestamp := time.Now().Unix();
	// then for userID, go review session handling

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