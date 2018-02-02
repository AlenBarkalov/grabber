package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine{
	r := gin.Default()

	// Три обработчика запросов
	r.GET("/",indexHandler)
	r.GET("/search/:category", collectHandler)
	r.GET("/result/:category", resultHandler)
	//r.Get("/",indexHandler)
	return r
}
