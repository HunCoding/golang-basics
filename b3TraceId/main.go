package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"log"
	"net/http"
	"time"
)

var tracer = otel.Tracer("B3-Tracer")

func initTracer() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("Exemplo-B3-Trace"),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(b3.New()))

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cleanup := initTracer()
	defer cleanup()

	go func() {
		http.HandleFunc("/trace", func(w http.ResponseWriter, r *http.Request) {
			ctx := otel.GetTextMapPropagator().Extract(
				r.Context(), propagation.HeaderCarrier(r.Header))

			ctx, span := tracer.Start(ctx, "ReceiveHandler")
			defer span.End()

			traceId := span.SpanContext().TraceID().String()
			fmt.Printf("TraceId gerado/recebido: %s \n", traceId)

			w.Header().Set("x-b3-traceid", traceId)
			w.Header().Set("b3-traceid", traceId)
		})

		fmt.Println("Começando execução do endpoint")
		err := http.ListenAndServe("localhost:8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Print("Aguardando endpoint inicializar")
	time.Sleep(1 * time.Second)

	client := &http.Client{}
	ctx, span := tracer.Start(context.Background(), "ClientRequest")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/trace", nil)
	if err != nil {
		log.Fatal(err)
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Printf("Resposta do servidor com header B3-TraceId: %s\n", resp.Header.Get("x-b3-traceid"))
}
