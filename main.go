package main

import (
	"fmt"

	"example.com/go-rest/db"
	"example.com/go-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world!")
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")
}
