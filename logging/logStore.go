package logging

import (
	"github.com/PMoneda/pruu/app"
	"github.com/gin-gonic/gin"
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
	msg := app.NewMessage(c)
	msg.ID = len(_map[key])
	_map[key] = append(_map[key], msg)
}

func Delete(key string) {
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

func FindAfter(key string, idx int) []app.Message {
	data, exist := _map[key]
	if !exist {
		return make([]app.Message, 0, 0)
	}
	return data[idx:]
}
