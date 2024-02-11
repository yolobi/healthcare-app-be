package middlewares

import (
	"bytes"
	"healthcare-capt-america/logger"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		args := make(map[string]interface{})
		if c.Request.Body != nil {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
			c.Request.Body = rdr
			args["request_body"] = c.Request.Body
			// args["method"] = c.Request.Method
			// args["status"] = c.Writer.Status()
			// args["uri"] = c.Request.RequestURI
		}

		logger.Log.WithFields(args).Info("Request")

		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		err := c.Errors.Last()
		if err == nil {
			logger.Log.Info("Success")
			return
		} else {
			logger.Log.Error("Error")
		}

		args = make(map[string]interface{})
		args["method"] = c.Request.Method
		args["uri"] = c.Request.RequestURI
		args["status"] = c.Writer.Status()
		args["latency"] = endTime.Sub(startTime).Seconds()

		logger.Log.WithFields(args).Error(err)

	}
}
