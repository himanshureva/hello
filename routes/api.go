package routes

import (
	"hello/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.POST("/todos", controllers.TodosCreate)
		api.GET("/todos", controllers.TodosIndex)
		api.GET("/todos/:id", controllers.TodosShow)
		api.PUT("/todos/:id", controllers.TodosUpdate)
		api.DELETE("/todos/:id", controllers.TodosDelete)
	}
}
