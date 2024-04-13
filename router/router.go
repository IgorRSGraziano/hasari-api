package router

import (
	"hasari-api/handlers/music"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.GET("/musics", music.List)
	r.Run(":8080")
}
