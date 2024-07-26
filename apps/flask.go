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
		filepath.Join(projectName, "static/css/style.css"): `
.disable-dbl-tap-zoom {
    touch-action: manipulation;
  }

  * { touch-action: manipulation; }
`,
		filepath.Join(projectName, "templates/header.html"): `
<!DOCTYPE html>
<html lang="pt-BR">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
    <link rel="stylesheet" href="../static/css/main.css">
    <link rel="stylesheet" href="../static/css/style.css">
    <link rel="manifest" href="../static/manifest.json">
    <link rel="stylesheet" href="/node_modules\flowbite\dist\flowbite.min.css">
    <script src="https://unpkg.com/hyperscript.org@0.9.12"></script>
    <meta name="theme-color" content="#ffffff">
    <link rel="icon" href="../static/img/favicon.ico" type="image/x-icon">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@100;200;300;400;500;600;700&display=swap"
        rel="stylesheet">

    <script src="../static/js/htmx.min.js"></script>
    <title>Maxcell O.S.</title>

</head>

<body class="font-[Poppins] bg-gradient-to-t from-slate-400 to-slate-200 min-h-screen flex flex-col">
    <header class="bg-white">
        <nav class="bg-white border-gray-200 ">
            <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                <a href="https://mxrota.shop/" class="flex items-center space-x-3 rtl:space-x-reverse">
                    <img src="/static/img/maxcell_celulares.svg" class="h-10" alt="Maxcell Logo" />
                    <span class="self-center text-2xl font-semibold whitespace-nowrap "></span>
                </a>
                <button data-collapse-toggle="navbar-default" type="button"
                    class="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 "
                    aria-controls="navbar-default" aria-expanded="false">
                    <span class="sr-only">Open main menu</span>
                    <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                        viewBox="0 0 17 14">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M1 1h15M1 7h15M1 13h15" />
                    </svg>
                </button>
                <div class="hidden w-full md:block md:w-auto" id="navbar-default">
                    <ul
                        class="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white ">

                        <li>
                            <a hx-get="/inicio" hx-target="#main" hx-trigger="click" hx-indicator="#spinner" href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Home</a>
                        </li>
                        <li>
                            <a hx-get="/aguardando" hx-target="#main" hx-trigger="click" hx-indicator="#spinner"
                                href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Aguardando</a>
                        </li>
                        <li>
                            <a hx-get="/iniciado" hx-target="#main" hx-trigger="click" hx-indicator="#spinner" href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Iniciada</a>
                        </li>
                        <li>
                            <a hx-get="/aguardandoaprovacao" hx-target="#main" hx-trigger="click"
                                hx-indicator="#spinner" href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Aguardando
                                Aprovação</a>
                        </li>
                        <li>
                            <a hx-get="/aguardandopeca" hx-target="#main" hx-trigger="click" hx-indicator="#spinner"
                                href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Aguardando
                                Peça</a>
                        </li>
                        <li>
                            <a hx-get="/senha" hx-target="#main" hx-trigger="click" hx-indicator="#spinner" href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Aguardando
                                Senha</a>
                        </li>
                        <li>
                            <a hx-get="/concluido" hx-target="#main" hx-trigger="click" hx-indicator="#spinner" href="#"
                                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-slate-400 md:p-0 ">Concluída</a>
                        </li>

                    </ul>
                </div>
            </div>
        </nav>
    </header>

`,
		filepath.Join(projectName, "templates/footer.html"): `

        <script src="../static/js/flowbite.min.js"></script>
        </body>
        <footer>
            <div class="footer bg-slate-700 flex justify-center items-center p-1">
                <p class="text-white text-sm">Sita & Sita | Copyright &copy; 2024</p>
            </div>
        </footer>
        
        </html>`,
		filepath.Join(projectName, "templates/index.html"): `
{% include 'header.html'%}
<main class="flex-grow">

</main>
<script>
    htmx.onLoad(function (content) {
        initFlowbite();
    })
</script>
{% include 'footer.html'%}
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
