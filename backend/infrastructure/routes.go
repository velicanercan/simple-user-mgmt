package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	return GinRouter{
		Gin: router,
	}

}
