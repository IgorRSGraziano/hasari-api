package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {})
	r.Run(":8080")
}
