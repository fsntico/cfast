package apps

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// FlaskCmd is the command to create a Flask project structure.
var FlaskCmd = &cobra.Command{
	Use:   "flask",
	Short: "Create a Flask project structure",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a project name")
			return
		}
		projectName := args[0]
		createFlaskProject(projectName)
	},
}

func createFlaskProject(projectName string) {
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
		filepath.Join(projectName, "main.py"): `from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'

if __name__ == '__main__':
    app.run(debug=True)
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
    <title>Projeto Flask</title>
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

	fmt.Printf("Flask project '%s' created successfully!\n", projectName)
}
