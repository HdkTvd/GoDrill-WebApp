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

Make a GET request to get list of all users by using cmd - `curl localhost:9090/GET/users`

Make a POST request to add users from csv file by using cmd - `curl localhost:9090/POST/users -XPOST -d "<.csv filepath>"`

OR

Use [Postman](https://www.postman.com/downloads/) for testing.
