package hub

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

var ginLogger = logrus.WithField("hub", "gin")

func ginRequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		ginLogger.Infof("| %3d | %13v | %15s | %s | %s |",
			c.Writer.Status(),
			time.Now().Sub(startTime),
			c.ClientIP(),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}
