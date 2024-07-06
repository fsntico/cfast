package apps

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var FiberCmd = &cobra.Command{
	Use:   "fiber",
	Short: "Create a Fiber project structure",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a project name")
			return
		}
		projectName := args[0]
		createFiberProject(projectName)
	},
}

func createFiberProject(projectName string) {
	dirs := []string{
		projectName,
		filepath.Join(projectName, "static"),
		filepath.Join(projectName, "templates"),
		filepath.Join(projectName, "static", "src"),
		filepath.Join(projectName, "static", "css"),
		filepath.Join(projectName, "static", "js"),
		filepath.Join(projectName, "static", "img"),
	}

	files := map[string]string{
		filepath.Join(projectName, "main.go"): `package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

		
		func main() {

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Static("/static", "./static")
	app.Get("/", Index)
	
	log.Fatal(app.Listen(":8520"))
	}
func Index(c *fiber.Ctx) error {

	return c.Render("bem_vindo", fiber.Map{})
}


`, filepath.Join(projectName, "static/src/input.css"): `@tailwind base;
@tailwind components;
@tailwind utilities;
`, filepath.Join(projectName, "tailwind.config.js"): `/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ["./templates/**/*.{html,htm}","./static/js/*.js"],
  theme: {
    extend: {
      colors : {
        primeira : '#0a100d',
        segunda : '#b9baa3',
        terceira : '#d6d5c9',
        quarta : '#a22c29',
        quinta : '#902923',
        
      },
      screens: {
        'hover-hover': {'raw': '(hover: hover)'
      },
    },
  },
},
}
`, filepath.Join(projectName, "package.json"): fmt.Sprintf(`
{
  "name": "%s",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "create-css": "npx tailwindcss -i ./static/src/input.css -o ./static/css/main.css --watch"
  },
  "author": "",
  "license": "ISC",
  "dependencies": {
    "flowbite": "^2.2.1",
    "tailwindcss": "^3.4.4"
  },
  "description": ""
}`, projectName), filepath.Join(projectName, "static/css/main.css"): ``,
		filepath.Join(projectName, "templates/index.html"): `<!DOCTYPE html>
<html lang="pt-BR">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/static\css\main.css">
    <link rel="icon" href="\static\img\favicon.ico" type="image/x-icon">
    <title>Projeto Fiber</title>
</head>

<body>

</body>

</html>
`,
	}

	// Create directories
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Create files
	for filePath, fileContent := range files {
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filePath, err)
			return
		}
		defer f.Close()

		_, err = f.WriteString(fileContent)
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			return
		}
	}

	fmt.Printf("Fiber project '%s' created successfully!\n", projectName)
}
