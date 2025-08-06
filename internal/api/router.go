package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/task", CreateTask)
	r.GET("/task/:id", GetTask)
	r.PUT("/task", UpdateTask)
	r.DELETE("/task/:id", DeleteTask)

	return r
}
