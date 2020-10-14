package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pfurtak/learning-gin/controller"
	"github.com/pfurtak/learning-gin/middleware"
	"github.com/pfurtak/learning-gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func logOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	logOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":5000")
}
