package handler

import (
	"gauntlet-be/m/v2/database"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	router.Use(CORSMiddleware())

	database.InitDb()

	v1 := router.Group("api/v1")

	v1.GET("/allocation-logs", GetAllocationLogs)
	v1.GET("/request-logs", GetRequestLogs)
	v1.GET("/redis-queue", GetNSELedhersLogs)
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
