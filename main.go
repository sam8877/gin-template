package main

import (
	"gin-template/midware"
	"gin-template/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), midware.Recovery())

	r.Use(midware.AuthMidware())
	r.POST("/create_token", service.CreateToken())
	r.POST("/refresh_token", service.RefreshToken())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
