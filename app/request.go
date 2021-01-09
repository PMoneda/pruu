package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"strings"
	"time"

	"io/ioutil"

	"net/http/httputil"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//Dump request structure
type Dump struct {
	CreatedAt time.Time `json:"created_at"`
	Tag       string    `json:"tag"`
	Checksum  string    `json:"body_checksum"`
	UID       string    `json:"id"`
	Opened    bool      `json:"opened"`
	Value     string    `json:"value"`
	Method    string    `json:"method"`
	BodySize  int64     `json:"body_size"`
	URI       string    `json:"uri"`
}

//Message stores a basic string struct
type Message struct {
	ID           int        `json:"id"`
	Key          string     `json:"key"`
	CreatedAt    time.Time  `json:"-"`
	CreatedAtStr string     `json:"created_at"`
	Value        string     `json:"value"`
	Level        string     `json:"level"`
	Tags         url.Values `json:"tags"`
	IsOlder      bool       `json:"is_older"`
}

func NewMessage(c *gin.Context) Message {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	stream := buf.String()
	level := strings.ToUpper(c.Query("level"))
	values := c.Request.URL.Query()
	values.Del("level")
	now := time.Now()
	nowStr := now.Format("2006-01-02T15:04:05")
	if err != nil {
		return Message{Key: uuid.NewV4().String(), CreatedAtStr: nowStr, CreatedAt: now, Value: err.Error()}
	}
	return Message{Key: uuid.NewV4().String(), CreatedAtStr: nowStr, Tags: values, Level: level, CreatedAt: now, Value: string(stream)}
}

//NewDump from request
func NewDump(c *gin.Context) Dump {
	full, _ := httputil.DumpRequest(c.Request, true)
	b := ioutil.NopCloser(c.Request.Body)
	body, _ := ioutil.ReadAll(b)

	dump := Dump{
		CreatedAt: time.Now(),
		Tag:       c.Param("tag"),
		Opened:    false,
		Checksum:  Sha256(body),
		UID:       uuid.NewV4().String(),
		Value:     string(full),
		Method:    c.Request.Method,
		BodySize:  c.Request.ContentLength,
		URI:       c.Request.RequestURI,
	}
	return dump
}

func Sha256(s []byte) string {
	h := sha256.New()
	h.Write(s)
	sEnc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sEnc
}
