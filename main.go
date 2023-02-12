package main

import (
	"log"
	"os"

	"github.com/adailsonm/desafio-sword/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := gin.Default()
	routes.AddRoutes(app)
	port := os.Getenv("PORT")
	log.Printf("API Listening on port %s", port)

	if err := app.Run(); err != nil {
		log.Fatalf("not start server: %v", err)
	}
}
