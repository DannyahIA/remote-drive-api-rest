package main

import (
	"fmt"

	"github.com/DannyahIA/personal-server/internal/router"
)

func main() {
	r := router.SetupRoutes()

	fmt.Println("Server is running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
