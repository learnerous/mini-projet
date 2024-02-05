package main

import (
	"fmt"
	routers "mini_projet/Routers"
)

func main() {
	r := routers.SetupRouter()
	fmt.Println("Server started on port 8080")
	r.Run()
}
