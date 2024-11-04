package main

import (
	"fmt"
	"log"

	"github.com/DannyahIA/personal-server/internal/database"
	"github.com/DannyahIA/personal-server/internal/router"
)

func main() {
	r := router.SetupRoutes()

	err := database.ConnectDatabase()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Server is running on port 8080")
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
