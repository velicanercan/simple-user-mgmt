package prometheus

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// RequestDurationMiddleware calculates the duration of the request
func RequestDurationMiddleware(reqDur *prometheus.HistogramVec) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		duration := endTime.Sub(startTime).Seconds()

		reqDur.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.Request.URL.Path).Observe(duration)
	}
}

// RequestCounterMiddleware increments the request counter
func RequestCounterMiddleware(reqCnt *prometheus.CounterVec) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		reqCnt.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.HandlerName(), c.Request.Host, c.Request.URL.Path).Inc()
	}
}

// RequestSizeMiddleware calculates the request size
func RequestSizeMiddleware(reqSz prometheus.Summary) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		reqSz.Observe(float64(c.Request.ContentLength))
	}
}

// ResponseSizeMiddleware calculates the response size
func ResponseSizeMiddleware(resSz prometheus.Summary) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		resSz.Observe(float64(c.Writer.Size()))
	}
}

// ResponseLatencyMiddleware calculates the latency of the response
func ResponseLatencyMiddleware(resLat *prometheus.HistogramVec) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		endTime := time.Now()
		duration := endTime.Sub(startTime).Seconds()

		resLat.WithLabelValues(strconv.Itoa(c.Writer.Status()), c.Request.Method, c.Request.URL.Path).Observe(duration)
	}
}
