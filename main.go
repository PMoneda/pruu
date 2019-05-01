package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/PMoneda/pruu/dump"
	"github.com/PMoneda/pruu/logging"
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
			"dumps": dump.FindByKey(c.Param("key")),
		})
	})

	router.GET("/log/:key", func(c *gin.Context) {
		c.HTML(http.StatusOK, "log.tmpl.html", gin.H{
			"logs": logging.FindByKey(c.Param("key")),
		})
	})

	router.DELETE("/dump/:key", func(c *gin.Context){
		mutex.Lock()
		defer mutex.Unlock()
		k := c.Param("key")
		dump.Delete(k, c)
		c.String(200, "OK")
	})

	router.POST("/dump/:key", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		k := c.Param("key")
		dump.Save(k, c)
		c.String(200, "OK")
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl.html", nil)
	})

	router.POST("/log/:key", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		k := c.Param("key")
		logging.Save(k, c)
		c.String(200, "OK")
	})
	

	router.Run(":" + port)
}
