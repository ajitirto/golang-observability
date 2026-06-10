package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP Requests",
		},
		[]string{"path"},
	)

  httpDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name: "http_request_duration_seconds",
        Help: "HTTP Duration",
    },
    []string{"path"},
)
)

func initMetrics() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(httpDuration)
}
