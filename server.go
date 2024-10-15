package main

import (
	"fmt"
	"main/ioapi"
	"os"
)

func main() {
	fmt.Println("Starting")

	swagger, err := ioapi.GetSwagger()
	if err != nil {
		fmt.Printf("Error loading swagger spec:\n%s", err)
		os.Exit(1)
	}
	fmt.Println(swagger)

}
