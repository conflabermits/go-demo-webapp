package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests",
	})
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request duration in seconds",
		Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

func New() *prometheus.Registry {
	return prometheus.NewRegistry()
}

func Middleware() func(http.Handler) http.Handler {
	return promhttp.InstrumentHandlerCounter(requestCount, promhttp.InstrumentHandlerDuration(requestDuration, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle request here
	})))
}
