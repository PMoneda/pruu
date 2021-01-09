package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/PMoneda/pruu/app"
	"github.com/PMoneda/pruu/dump"
	"github.com/PMoneda/pruu/logging"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var mutex sync.Mutex
var _channelMap map[string]*DataReceiver
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type DataReceiver struct {
	Connections []*websocket.Conn
	Channel     chan app.Message
}

func remove(s []*websocket.Conn, i int) []*websocket.Conn {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func distributeMessage(key string) {
	_, exist := _channelMap[key]
	if !exist {
		return
	}
	for msg := range _channelMap[key].Channel {
		mutex.Lock()
		jj := _channelMap[key]
		conns := &jj.Connections
		for i := 0; i < len((*conns)); i++ {
			err := (*conns)[i].WriteJSON(msg)
			if err != nil {
				*conns = append((*conns)[:i], (*conns)[i+1:]...)
				i = i - 1
				if i < 0 {
					i = 0
				}

			}
		}
		mutex.Unlock()
	}

}
func wshandler(c *gin.Context, r *http.Request) {
	conn, err := wsupgrader.Upgrade(c.Writer, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	key := c.Param("key")
	msgs := logging.FindByKey(key)
	_, exist := _channelMap[key]
	if !exist {
		_channelMap[key] = &DataReceiver{
			Connections: make([]*websocket.Conn, 1),
			Channel:     make(chan app.Message),
		}
		_channelMap[key].Connections[0] = conn
		go distributeMessage(key)
		for _, msg := range msgs {
			_channelMap[key].Channel <- msg
		}
	} else {
		mutex.Lock()
		hub := _channelMap[key]
		hub.Connections = append(hub.Connections, conn)
		mutex.Unlock()
	}
}

func main() {
	_channelMap = make(map[string]*DataReceiver)
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
		msg := logging.Save(k, c)
		ch, exist := _channelMap[k]
		if exist {
			ch.Channel <- msg
		}
		c.String(200, "OK")
	})

	router.GET("/ws/log/:key", func(c *gin.Context) {
		wshandler(c, c.Request)
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
