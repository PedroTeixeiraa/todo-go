package main

import (
	"fmt"
	"github.com/PedroTeixeiraa/todo-go/config"
	"github.com/PedroTeixeiraa/todo-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application...")

	config.ConnectDB()
	fmt.Println("Database connection successful!")

	r := gin.Default()
	routes.SetupRoutes(r)
	fmt.Println("Routes configured. Server ready to run.")

	r.Run(":8080")
}
