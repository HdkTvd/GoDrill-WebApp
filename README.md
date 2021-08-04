# GoDrill-WebApp

## A Project in GoLang.

### Prerequisite programs
- Golang
- MySQL database

### How to execute this project?

Clone this repo -
`git clone https://github.com/HdkTvd/GoDrill-WebApp.git`

Go to your project directory -
`cd ./GoDrill-WebApp`

Install Dependencies -
`go mod tidy`

In your project directory run command -
`go run main.go`

### How to test this project?

Make a POST request to localhost:9090/users as - `curl localhost:9090/users -XPOST -d "<.csv filepath>"`
Make a GET request to localhost:9090/ as - `curl localhost:9090/`

OR

Use [Postman](https://www.postman.com/downloads/) for testing.
