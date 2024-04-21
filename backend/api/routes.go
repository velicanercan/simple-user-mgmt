package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/velicanercan/simple-user-mgmt/prometheus"
)

type GinRouter struct {
	Gin *gin.Engine
}

// NewGinRouter returns a new gin router
func NewGinRouter() GinRouter {
	router := gin.New()

	// Add Prometheus metrics
	prometheus.UseMetrics(router)
	prometheus.RegisterMetrics(router)
	
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "I'm alive!"})
	})

	router.Use(LoggerMiddleware())
	router.Use(CORSMiddleware())
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

// CORSMiddleware is a middleware that handles CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_POLICY"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
