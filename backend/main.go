package main

import (
	"log"
	"online_library/controllers"
	db "online_library/databes"
	routes "online_library/routers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading environment file: ", err)
	}
	db := db.ConnectDB()
	controllers.SetCollection(db.Collection("Books"))
	controllers.SetUserCollection(db.Collection("Users"))
	app := gin.Default()

	app.RedirectTrailingSlash = false
	// Use Gin CORS middleware with the desired config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},            // Allow specific origin
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},   // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Allowed headers
		AllowCredentials: true,                                         // Allow credentials like cookies, etc.
	}))
	app.Use(gin.Logger())
	routes.SetupBookRoutes(app)
	routes.SetupUserRoutes(app)
	app.SetTrustedProxies([]string{"127.0.0.1"}) // Set trusted proxies as needed

	port := os.Getenv("PORT")
	if port == "" {
		port = "7788"
	}
	log.Fatal(app.Run(":" + port))
}
