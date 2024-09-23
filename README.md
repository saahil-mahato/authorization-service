## Authentication Service

This project is a micro service for RBAC. It uses the permify api.

[Link to Permify Docs](https://docs.permify.co/permify-overview/intro)


## How to start the Application in local

1. Download and install the Go programming language. You can follow the instructions [here](https://go.dev/doc/install)

2. Setup Docker for your machine. You can find the instructions [here](https://docs.docker.com/engine/install/)

3. Run the Permify service with the following command

`docker run -p 3476:3476 -p 3478:3478  ghcr.io/permify/permify`

**Note**: You might require admin permissions so use `sudo` command if permission is denied.

4. Navigate to the project directory and install the dependencies using the following commands:\n
`go mod download`\n
`go mod tidy`

5. Run the application with the command `go run main.go`.

**Note**: You might also want to setup a debugger for development. If you are using vscode, please follow the instructions [here](https://github.com/golang/vscode-go/wiki/debugging)