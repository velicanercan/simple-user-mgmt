package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GinRouter struct {
	Gin *gin.Engine
}

// NewGinRouter returns a new gin router
func NewGinRouter() GinRouter {
	router := gin.New()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "I'm alive!"})
	})

	router.Use(LoggerMiddleware())
	return GinRouter{
		Gin: router,
	}

}

// LoggerMiddleware logs the incoming requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Infof("Request received: %s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()
	}
}
