package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rpuglielli/person-api/internal/infrastructure/http/handler"
)

func SetupRouter(personHandler *handler.PersonHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	persons := r.Group("/persons")
	{
		persons.GET("", personHandler.ListPersons)
		persons.POST("", personHandler.CreatePerson)
		persons.GET("/:id", personHandler.GetPerson)
		persons.PUT("/:id", personHandler.UpdatePerson)
		persons.DELETE("/:id", personHandler.DeletePerson)
	}

	return r
}
