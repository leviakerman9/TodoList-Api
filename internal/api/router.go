package api

import "github.com/gin-gonic/gin"

func SetupRouter(handler *APIHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/task", handler.CreateTask)
	r.GET("/task/:id", handler.GetTask)
	r.PUT("/task", handler.UpdateTask)
	r.DELETE("/task/:id", handler.DeleteTask)

	return r
}
