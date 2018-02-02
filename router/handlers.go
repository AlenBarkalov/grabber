package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"grabber/news"
)

func indexHandler(c *gin.Context){
	c.String(http.StatusOK, "Привет МИР!")
}

func collectHandler(c *gin.Context){
	//c.String(http.StatusOK, "Привет МИР!")
	category := c.Param("category")

	news.Collect(category)
	c.String(http.StatusOK,"Инициирован поиск")
}

func resultHandler(c *gin.Context){
	//c.String(http.StatusOK, "Привет МИР!")
	category := c.Param("category")

	topics := news.Result(category)
	c.JSON(http.StatusOK, topics)
}