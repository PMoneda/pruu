package main

import (
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/PMoneda/pruu/app"
	"github.com/PMoneda/pruu/dump"
	"github.com/PMoneda/pruu/logging"
	"github.com/gin-gonic/gin"
)

var mutex sync.Mutex

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.PATCH("/log/:key", func(c *gin.Context) {

		k := c.Param("key")
		list := logging.FindByKey(k)
		if len(list) > 0 {
			c.JSON(200, list[len(list)-1])
		} else {
			c.String(404, "not found")
		}

	})

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

	router.GET("/json/dump/:key", func(c *gin.Context) {
		c.JSON(http.StatusOK, dump.FindByKey(c.Param("key")))
	})

	router.GET("/log/:key", func(c *gin.Context) {
		data := getMessage(c)
		c.HTML(http.StatusOK, "log.tmpl.html", gin.H{
			"logs": data,
		})
	})

	router.GET("/json/log/:key", func(c *gin.Context) {
		data := getMessage(c)
		c.JSON(http.StatusOK, data)
	})

	router.DELETE("/dump/:key", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		k := c.Param("key")
		dump.Delete(k)
		c.String(200, "OK")
	})

	router.DELETE("/log/:key", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		k := c.Param("key")
		logging.Delete(k)
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

func getMessage(c *gin.Context) []app.Message {
	var data []app.Message
	if c.Query("offset") != "" {
		idx, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.String(500, err.Error())
			return nil
		}
		data = logging.FindAfter(c.Param("key"), idx)
	} else {
		data = logging.FindByKey(c.Param("key"))
	}
	return data
}
