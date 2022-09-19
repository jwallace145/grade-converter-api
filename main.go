package main

/*
IMPORTS
*/
import (
	"example/grade-converter-api/api"
)

/*
CONSTANTS
*/
const HOST string = "localhost"
const PORT int = 8080

/*
MAIN
*/
func main() {

	// start grade converter api
	api.New(HOST, PORT)
}
