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
`, filepath.Join(projectName, "static/css/style.css"): `
.disable-dbl-tap-zoom {
    touch-action: manipulation;
  }

  * { touch-action: manipulation; }
`,
		filepath.Join(projectName, "tailwind.config.js"): `/** @type {import('tailwindcss').Config} */
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
		filepath.Join(projectName, "templates/header.html"): `
{{define "header"}}
<!DOCTYPE html>
<html lang="pt-BR">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
    <link rel="stylesheet" href="../static/css/main.css">
    <link rel="stylesheet" href="../static/css/style.css">
    <link rel="icon" href="../static/img/favicon.ico" type="image/x-icon">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@100;200;300;400;500;600;700&display=swap"
        rel="stylesheet">
    <script type="module" src="https://unpkg.com/ionicons@4.5.10-0/dist/ionicons/ionicons.esm.js"></script>
    <script nomodule="" src="https://unpkg.com/ionicons@4.5.10-0/dist/ionicons/ionicons.js"></script>
    <title>Projeto Fiber</title>
</head>

<body class="font-[Poppins] bg-gradient-to-t from-slate-400 to-slate-200 h-screen flex flex-col">
    <header class="bg-white">
        <nav class="flex justify-between items-center w-[92%]  mx-auto">
            <div>
                <img class="w-16 cursor-pointer" src="https://cdn-icons-png.flaticon.com/512/5968/5968204.png"
                    alt="...">
            </div>
            <div
                class="nav-links duration-500 md:static absolute bg-white md:min-h-fit min-h-[60vh] left-0 top-[-100%] md:w-auto  w-full flex items-center px-5">
                <ul class="flex md:flex-row flex-col md:items-center md:gap-[4vw] gap-8">
                    <li>
                        <a class="hover:text-gray-500" href="#">Home</a>
                    </li>
                    <li>
                        <a class="hover:text-gray-500" href="#">
                            <div class="flex items-center gap-2">
                                Entregas
                            </div>
                        </a>
                    </li>
                    <li>
                        <a class="hover:text-gray-500" href="#">
                            <div class="flex items-center gap-2">
                                Separado
                            </div>
                        </a>
                    </li>
                    <li>
                        <a class="hover:text-gray-500" href="#">
                            <div class="flex items-center gap-2">Em Rota</div>
                        </a>
                    </li>
                    <li>
                        <a class="hover:text-gray-500" href="#">
                            <div class="flex items-center gap-2">Entregue</div>
                        </a>
                    </li>
                </ul>
            </div>
            <div class="flex items-center gap-6">
                <!-- <button class="bg-[#a6c1ee] text-white px-5 py-2 rounded-full hover:bg-[#87acec]">Sign in</button> -->
                <ion-icon onclick="onToggleMenu(this)" name="menu" class="text-3xl cursor-pointer md:hidden"></ion-icon>
            </div>
    </header>
    {{end}}
`,
		filepath.Join(projectName, "templates/footer.html"): `
{{define "footer"}}
<script>
    const navLinks = document.querySelector('.nav-links')
    function onToggleMenu(e) {
        e.name = e.name === 'menu' ? 'close' : 'menu'
        navLinks.classList.toggle('top-[9%]')
    }
</script>

</body>
<footer>
    <div class="footer bg-slate-700 flex justify-center items-center p-1">
        <p class="text-white text-sm">Sita & Sita | Copyright &copy; 2024</p>
    </div>
</footer>
</html>
{{end}}`,
		filepath.Join(projectName, "templates/index.html"): `
    {{template "header"}}
<main class="flex-grow">

</main>
{{template "footer"}}
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
