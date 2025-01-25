package main

import (
	"fmt"
	"github.com/PedroTeixeiraa/todo-go/config"
	"github.com/PedroTeixeiraa/todo-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Iniciando aplicação...")

	config.ConnectDB()
	fmt.Println("Conexão ao banco bem-sucedida!")

	r := gin.Default()
	routes.SetupRoutes(r)
	fmt.Println("Rotas configuradas. Servidor pronto para rodar.")

	r.Run(":8080")
}
