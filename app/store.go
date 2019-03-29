package app

import (
	"github.com/gin-gonic/gin"
)

var _map map[string][]Dump

func init() {
	_map = make(map[string][]Dump)
}
func Save(key string, c *gin.Context) {
	_, exist := _map[key]
	if !exist {
		_map[key] = make([]Dump, 0, 0)
	}
	_map[key] = append([]Dump{NewDump(c)} ,_map[key]...)
}

func Delete(key string, c *gin.Context) {
	_, exist := _map[key]
	if exist {
		_map[key] = make([]Dump, 0, 0)
	}
}

func FindByKey(key string) []Dump {
	data, exist := _map[key]
	if !exist {
		return make([]Dump, 0, 0)
	}
	return data
}
