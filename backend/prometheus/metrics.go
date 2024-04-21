package prometheus

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

// RegisterMetrics registers the prometheus metrics to the router
func RegisterMetrics(r *gin.Engine) {
	reqDur := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "The HTTP request latencies in seconds.",
		Buckets: prometheus.DefBuckets,
	}, []string{"code", "method", "url"})

	reqCnt := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	}, []string{"code", "method", "handler", "host", "url"})

	resSz := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "response_size_bytes",
		Help: "The HTTP response sizes in bytes.",
	})

	reqSz := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "request_size_bytes",
		Help: "The HTTP request sizes in bytes.",
	})

	resLat := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "response_latency_seconds",
		Help:    "The HTTP response latencies in seconds.",
		Buckets: prometheus.DefBuckets,
	}, []string{"code", "method", "url"})

	prometheus.MustRegister(reqDur)
	prometheus.MustRegister(reqCnt)
	prometheus.MustRegister(resSz)
	prometheus.MustRegister(reqSz)
	prometheus.MustRegister(resLat)

	r.Use(RequestDurationMiddleware(reqDur))
	r.Use(RequestCounterMiddleware(reqCnt))
	r.Use(ResponseSizeMiddleware(resSz))
	r.Use(RequestSizeMiddleware(reqSz))
	r.Use(ResponseLatencyMiddleware(resLat))
}

// UseMetrics adds prometheus metrics to the router
func UseMetrics(router *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin", CreateMetrics())

	p.Use(router)
}

// CreateMetrics creates the prometheus metrics
func CreateMetrics() []*ginprometheus.Metric {
	var metrics []*ginprometheus.Metric

	metrics = append(metrics, RequestCounter())
	metrics = append(metrics, RequestDuration())
	metrics = append(metrics, ResponseSize())
	metrics = append(metrics, RequestSize())
	metrics = append(metrics, ResponseLatency())

	return metrics
}

// Request Counter metric
func RequestCounter() *ginprometheus.Metric {
	var reqCnt = &ginprometheus.Metric{
		ID:          "reqCnt",
		Name:        "requests_total",
		Description: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		Type:        "counter_vec",
		Args:        []string{"code", "method", "handler", "host", "url"},
	}

	return reqCnt
}

// Request Duration metric
func RequestDuration() *ginprometheus.Metric {
	var reqDur = &ginprometheus.Metric{
		ID:          "reqDur",
		Name:        "request_duration_seconds",
		Description: "The HTTP request latencies in seconds.",
		Type:        "histogram_vec",
		Args:        []string{"code", "method", "url"},
	}

	return reqDur
}

// Request Size metric
func ResponseSize() *ginprometheus.Metric {
	var resSz = &ginprometheus.Metric{
		ID:          "resSz",
		Name:        "response_size_bytes",
		Description: "The HTTP response sizes in bytes.",
		Type:        "summary",
	}

	return resSz
}

// Response Size metric
func RequestSize() *ginprometheus.Metric {
	var reqSz = &ginprometheus.Metric{
		ID:          "reqSz",
		Name:        "request_size_bytes",
		Description: "The HTTP request sizes in bytes.",
		Type:        "summary",
	}

	return reqSz
}

// Response Latency metric
func ResponseLatency() *ginprometheus.Metric {
	var resLat = &ginprometheus.Metric{
		ID:          "resLat",
		Name:        "response_latency_seconds",
		Description: "The HTTP response latencies in seconds.",
		Type:        "summary",
	}

	return resLat
}
