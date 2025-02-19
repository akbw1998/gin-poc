package main

import (
	"io"
	"net/http"
	"os"
	"practice/controllers"
	"practice/entities"
	"practice/middlewares"
	"practice/services"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    = services.New()
	videoController = controllers.New(videoService)
)

func setupLoggerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
}

func main() {
	setupLoggerOutput()

	server := gin.New()
	server.Use(middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/healthz", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK",
			})
		})

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			response := videoController.Save(ctx)

			switch res := response.(type) {
			case error:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": res.Error(),
				})
			case entities.Video:
				ctx.JSON(http.StatusOK, res)
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
