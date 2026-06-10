package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
}

func WithTrace(ctx context.Context, fields log.Fields) *log.Entry {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()

	if sc.HasTraceID() {
		fields["trace_id"] = sc.TraceID().String()
		fields["span_id"] = sc.SpanID().String()
	}

	return log.WithFields(fields)
}
