package routes

import (
	"online_library/controllers"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(app *gin.Engine) {
	app.GET("/api/book/", controllers.GetAllBooks)
	app.POST("/api/book/", controllers.CreateBook)
	app.GET("/api/book/:id", controllers.GetBook)
	app.PATCH("/api/book/:id", controllers.UpdateBook)
	app.DELETE("/api/book/:id", controllers.DeleteBook)
}
