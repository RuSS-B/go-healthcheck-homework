package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of requests.",
	},
	[]string{"path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path", "statusCode"})

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		//Just an anchor to slow down the response
		rndVal := rand.Intn(100)
		time.Sleep(time.Duration(rndVal) * time.Millisecond)

		label := r.RequestURI

		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(label).Inc()

		httpDuration.WithLabelValues(label, strconv.Itoa(statusCode)).Observe(time.Since(begin).Seconds())
	})
}

func rare500(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rndVal := rand.Intn(10)

		if rndVal == 0 {
			w.WriteHeader(500)
			_, _ = w.Write([]byte("This is a fake 500 Error!"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func initMetrics() {
	_ = prometheus.Register(totalRequests)
	_ = prometheus.Register(responseStatus)
	_ = prometheus.Register(httpDuration)
}
