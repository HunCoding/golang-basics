package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"time"
)

var tracer = otel.Tracer("Example-Go-Trace")

func initTracer() func() {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		log.Fatal(err)
	}

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("Example-Trace"))

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resources))
	otel.SetTracerProvider(traceProvider)

	return func() {
		err := traceProvider.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cleanup := initTracer()
	defer cleanup()

	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	doWork(ctx)
}

func doWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doWork")
	defer span.End()

	time.Sleep(100 * time.Millisecond)
	doSubWork(ctx)
}

func doSubWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doSubWork")
	defer span.End()

	time.Sleep(500 * time.Millisecond)
	doSubSubWork(ctx)
}

func doSubSubWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doSubSubWork")
	defer span.End()

	time.Sleep(1 * time.Second)

	fmt.Println("Codigo finalizado no ultimo metodo")
}
