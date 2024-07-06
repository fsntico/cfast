# cfast - CLI Application for Flask and Fiber Project Structure
## English
### Overview

cfast is a simple CLI application that generates a project structure for Flask or Fiber frameworks. It helps developers quickly set up a new project by creating necessary folders and files.
Installation

#### Clone the repository or Download binaries:

    git clone https://github.com/fsntico/cfast.git
    cd cfast

Build the application:

    go build -o cfast

Move the executable to a directory in your PATH (optional):

    mv cfast /usr/local/bin/

Usage

Run the cfast command followed by the framework name to generate the project structure.
Flask

    cfast flask <Project Name>

This command will cfast the following structure:

    flask

        project_name/
        │
        ├── main.py
        ├── static/
        │   ├── src/
        │       └── input.css
        │   └── images/
        ├── templates/
        └── README.md

Fiber

    
    cfast fiber
    

This command will cfast the following structure:

css

    project_name/
    │
    ├── main.go
    ├── static/
    │   ├── src/
    │       └── input.css
    │   └── images/
    ├── templates/
    └── README.md

Contributing

Feel free to open issues or submit pull requests with improvements.
License

This project is licensed under the MIT License.


# Português
## Visão Geral

cfast é uma aplicação CLI simples que gera a estrutura de projeto para os frameworks Flask ou Fiber. Ele ajuda os desenvolvedores a configurar rapidamente um novo projeto, criando as pastas e arquivos necessários.
Instalação

#### Clone o repositório ou faça download do binário:



    git clone https://github.com/seunomeusuario/cfast.git
    cd cfast

Compile a aplicação:


    go build -o cfast

Mova o executável para um diretório no seu PATH (opcional):

sh

    mv cfast /usr/local/bin/

Uso

Execute o comando cfast seguido do nome do framework para gerar a estrutura do projeto.
Flask


    cfast flask

Este comando criará a seguinte estrutura:

css

    project_name/
    │
    ├── main.py
    ├── static/
    │   ├── src/
    │       └── input.css
    │   └── images/
    ├── templates/
    └── README.md

Fiber


    cfast fiber <nome_do_projeto>

Este comando criará a seguinte estrutura:

css

    project_name/
    │
    ├── main.go
    ├── static/
    │   ├── src/
    │       └── input.css
    │   └── images/
    ├── templates/
    └── README.md

Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests com melhorias.
Licença

Este projeto está licenciado sob a Licença MIT.