package dump

import (
	"github.com/PMoneda/pruu/app"
	"github.com/gin-gonic/gin"
)

var _map map[string][]app.Dump

func init() {
	_map = make(map[string][]app.Dump)
}
func Save(key string, c *gin.Context) {
	_, exist := _map[key]
	if !exist {
		_map[key] = make([]app.Dump, 0, 0)
	}
	_map[key] = append([]app.Dump{app.NewDump(c)}, _map[key]...)
}

func Delete(key string) {
	_, exist := _map[key]
	if exist {
		_map[key] = make([]app.Dump, 0, 0)
	}
}

func FindByKey(key string) []app.Dump {
	data, exist := _map[key]
	if !exist {
		return make([]app.Dump, 0, 0)
	}
	return data
}
