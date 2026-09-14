## 1. Signals (The "What")

OpenTelemetry organizes observability into three primary pillars, called **Signals**:

- **Traces:** Track the progression of a single request as it moves through various services and functions.
    
- **Metrics:** Record aggregations over time (e.g., error rates, CPU usage, request counts).
    
- **Logs:** Record discrete events (OTel's logging signal is still maturing in Go, so Traces and Metrics are the most common starting points).
    

## 2. Providers (The "Engine")

Providers are the entry points for the OpenTelemetry SDK. They hold the configuration for how telemetry is processed and exported. You typically initialize these once at the start of your Go application.

- **`TracerProvider`:** Creates `Tracer` instances.
    
- **`MeterProvider`:** Creates `Meter` instances.
    

_Note: In Go, you often set a global provider using `otel.SetTracerProvider()`, which allows any package in your app to generate telemetry without passing the provider around._

## 3. Tracers, Meters, and Instruments (The "Tools")

You don't create telemetry directly from the Provider. Instead, you ask the Provider for a scoped tool, usually named after the Go package using it.

- **`Tracer`:** Created by a `TracerProvider`. The tracer is responsible for creating **Spans**.
    
- **`Meter`:** Created by a `MeterProvider`. The meter creates **Instruments** (like counters, gauges, and histograms) to record metrics.
    

## 4. Spans (The "Unit of Work")

A **Span** represents a single operation within a trace (e.g., a database query, an HTTP request, or a specific function execution).

- It contains a name, start/end timestamps, and contextual metadata.
    
- You can attach **Attributes** (key-value pairs), **Events** (timestamped logs within the span) to provide more detail and **errors**.
    
- In Go, a span is created and typically ended using `defer span.End()`.

```go
ctx, span := tracer.Start(ctx, "ProcessOrder")

defer span.End()

// 1. Add a standard event
span.AddEvent("validating_order", trace.WithAttributes(
	attribute.String("order.id", orderID),
))

// Simulate some business logic that might fail
err := performDatabaseUpdate(orderID)

if err != nil {
	// 2. Record the error details on the span
	span.RecordError(err, trace.WithAttributes(
		attribute.String("db.operation", "update"),
	))
	
	// 3. Mark the span status as Failed
	
	// Without this, the span will still appear as "Success" in your UI
	span.SetStatus(codes.Error, "Database update failed during order processing")
	
	return fmt.Errorf("failed to process order: %w", err)

}

// If successful, you can explicitly set Ok (though Unset/Success is usually the default)

span.SetStatus(codes.Ok, "Order processed successfully")
```
    
## 5. Instruments, Counters vs. Gauges vs. Histograms

These are the three primary types of **Instruments**. You choose which one to use based on the _shape_ of the data you want to record.

| **Instrument** | **What it does**                                                                      | **When to use it**                                                                            | **Example**                                             |
| -------------- | ------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| **Counter**    | A value that only ever goes up (monotonic).                                           | Counting discrete occurrences or events.                                                      | Total HTTP requests, total errors, bytes processed.     |
| **Gauge**      | A value that captures the current state (can go up or down).                          | Taking a "snapshot" of a current measurement.                                                 | Active web sockets, CPU temperature, available memory.  |
| **Histogram**  | Groups multiple recorded values into "buckets" to calculate statistical distribution. | Measuring rates, durations, or sizes to find averages, percentiles (e.g., P99), min, and max. | Request latency, payload size, database query duration. |

```go
package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// 1. Ask the global MeterProvider for a Meter named after your service/package
var meter = otel.Meter("my-billing-service")

func ProcessMetricsExample(ctx context.Context) {
	// ==========================================
	// 1. COUNTER
	// ==========================================
	// Create the instrument once (usually done at package initialization or struct creation)
	requestsCounter, _ := meter.Int64Counter(
		"http.requests.total",
		metric.WithDescription("Total number of HTTP requests processed"),
	)

	// Usage: Add values to it as events happen. It only goes up.
	requestsCounter.Add(ctx, 1)


	// ==========================================
	// 2. GAUGE
	// ==========================================
	// Create the instrument
	activeConnections, _ := meter.Int64Gauge(
		"db.connections.active",
		metric.WithDescription("Current number of active database connections"),
	)

	// Usage: Record the absolute current value. 
	// If you have 5 connections now, you record 5. If it drops to 2 later, you record 2.
	activeConnections.Record(ctx, 5)


	// ==========================================
	// 3. HISTOGRAM
	// ==========================================
	// Create the instrument
	requestLatency, _ := meter.Float64Histogram(
		"http.request.duration",
		metric.WithDescription("Duration of HTTP requests"),
		metric.WithUnit("s"), // explicitly stating the unit is a good practice
	)

	// Usage: Record a specific duration. The backend will aggregate this 
	// into percentiles (like "99% of requests took less than 0.25 seconds").
	requestLatency.Record(ctx, 0.145) // e.g., representing 145 milliseconds
}
```

## 6. Context and Propagators (The "Glue")

This is arguably the most important abstraction for Go developers. OpenTelemetry uses Go's `context.Context` to carry the "Span Context" (the trace ID and span ID) across function calls.

- When you start a new span, you pass in a `Context`, and the Tracer returns a _new_ `Context` containing the new span. You must pass this new context to downstream functions.
    
- **Propagators:** When your Go service makes an external call (like an HTTP request to another microservice), a Propagator injects this context into the HTTP headers so the next service can continue the same trace.
    

## 7. Exporters and Processors (The "Delivery")

These sit at the end of the SDK pipeline.

- **Processors:** Determine _how_ data is batched and sent (e.g., `BatchSpanProcessor` groups spans together before sending them to save network overhead).
    
- **Exporters:** Translate OTel data into a specific format and send it to a backend. The most common is the **OTLP Exporter** (OpenTelemetry Protocol), which sends data to an OpenTelemetry Collector or directly to backends like Jaeger, Datadog, or Honeycomb.