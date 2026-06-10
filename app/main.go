package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	initMetrics()
	initLogger()

	cleanup := initTracer()
	defer cleanup()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", homeHandler)

	log.Info("server started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	httpRequests.
		WithLabelValues(r.URL.Path).
		Inc()

	tracer := otel.Tracer("go-observability")

	ctx, span := tracer.Start(r.Context(), "home-request")
	defer span.End()

	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.path", r.URL.Path),
		attribute.String("user_agent", r.UserAgent()),
	)

	defer func() {
		duration := time.Since(start).Seconds()

		httpDuration.
			WithLabelValues(r.URL.Path).
			Observe(duration)

		log.WithFields(log.Fields{
			"path":       r.URL.Path,
			"method":     r.Method,
			"latency_ms": time.Since(start).Milliseconds(),
			"trace":      span.SpanContext().TraceID().String(),
		}).Info("request completed")
	}()

	_ = ctx

	fmt.Fprintf(w, "Hello Observability Stack 🚀")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.
		WithLabelValues("/health").
		Inc()

	fmt.Fprintf(w, "OK")
}
