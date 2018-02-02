package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"grabber/news"
)

// возрат строки пользователю
func indexHandler(c *gin.Context){
	c.String(http.StatusOK, "Привет МИР!")
}

// осуществляет поиск
func collectHandler(c *gin.Context){
	//c.String(http.StatusOK, "Привет МИР!")
	category := c.Param("category")

	news.Collect(category)
	c.String(http.StatusOK,"Инициирован поиск")
}


// получает результаты и возвращает в JSON
func resultHandler(c *gin.Context){
	//c.String(http.StatusOK, "Привет МИР!")
	category := c.Param("category")

	topics := news.Result(category)
	c.JSON(http.StatusOK, topics)
}