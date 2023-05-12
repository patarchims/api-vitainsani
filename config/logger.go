package config

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()

	startTime := time.Now()
	duration := time.Since(startTime)

	logger := log.Info()

	logger.Str("protocol", "http").
		Str("method", c.Request.Method).
		Str("path", c.Request.RequestURI).
		Dur("duration", duration).
		Msg("receive a HTTP request")

	if statusCode >= 400 {
		fmt.Println("Response body: " + blw.body.String())
	}
}
