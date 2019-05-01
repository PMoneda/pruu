package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/PMoneda/pruu/app"
)

 

var _map map[string][]app.Message


func init() {
	_map = make(map[string][]app.Message)
}
func Save(key string, c *gin.Context) {
	_, exist := _map[key]
	if !exist {
		_map[key] = make([]app.Message, 0, 0)
	}
	_map[key] = append(_map[key],app.NewMessage(c))
}

func Delete(key string, c *gin.Context) {
	_, exist := _map[key]
	if exist {
		_map[key] = make([]app.Message, 0, 0)
	}
}

func FindByKey(key string) []app.Message {
	data, exist := _map[key]
	if !exist {
		return make([]app.Message, 0, 0)
	}
	return data
}