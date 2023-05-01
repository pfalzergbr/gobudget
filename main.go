package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB running!")


	apiServer := NewAPIServer(":8080")
	apiServer.Run()
}
