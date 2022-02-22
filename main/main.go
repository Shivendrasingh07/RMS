package main

import (
	"fmt"
	"github.com/RMS/database"
	"github.com/RMS/routes"
)

func main() {

	err := database.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the Database successfully")

	server := routes.Route()

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
