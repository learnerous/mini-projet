package main

import (
	"fmt"

	"github.com/path/to/Routers" // Import the package that contains the SetupRouter function
)

func main() {
	r := Routers.SetupRouter()
	fmt.Println("Server started on port 8080")
	r.Run()
}
