package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/PMoneda/pruu/app"
	"github.com/gin-gonic/gin"
)

var mutex sync.Mutex

func main() {
	port := os.Getenv("PORT")
	if(port == ""){
		port = "8080"
	}
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/dump/:key", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"dumps": app.FindByKey(c.Param("key")),
		})
	})

	router.POST("/dump/:key", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		k := c.Param("key")
		app.Save(k, c)
		c.String(200, "OK")
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl.html", nil)
	})

	router.Run(":" + port)
}
