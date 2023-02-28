package handler

import (
	"gauntlet-be/m/v2/database"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	database.InitDb()

	v1 := router.Group("api/v1")

	v1.GET("/allocation-logs", GetAllocationLogs)
	v1.GET("/request-logs", GetRequestLogs)
	// v1.GET("/redis-queue", RedisQueue)

	return router
}
