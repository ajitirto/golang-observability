package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

func initTracer() func() {
	conn, _ := grpc.Dial("tempo:4317", grpc.WithInsecure())

	exporter, _ := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithGRPCConn(conn),
	)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tp)

	return func() {
		tp.Shutdown(context.Background())
	}
}
