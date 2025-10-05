package internal

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type Telemetry struct {
	RequestCounter  metric.Int64Counter
	RequestDuration metric.Float64Histogram
	MeterProvider   *sdkmetric.MeterProvider
}

func InitTelemetry(serviceName string) (*Telemetry, error) {
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(semconv.ServiceName(serviceName)),
	)
	if err != nil {
		return nil, err
	}

	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(exporter),
	)

	meter := meterProvider.Meter(serviceName)

	requestCounter, err := meter.Int64Counter("http_requests_total")
	if err != nil {
		return nil, err
	}

	requestDuration, err := meter.Float64Histogram("http_request_duration_seconds")
	if err != nil {
		return nil, err
	}

	return &Telemetry{
		RequestCounter:  requestCounter,
		RequestDuration: requestDuration,
		MeterProvider:   meterProvider,
	}, nil
}

func (t *Telemetry) RecordRequest(method, path string, statusCode int, duration time.Duration) {
	attrs := metric.WithAttributes(
		attribute.String("method", method),
		attribute.String("path", path),
		attribute.Int("status_code", statusCode),
	)

	t.RequestCounter.Add(context.Background(), 1, attrs)
	t.RequestDuration.Record(context.Background(), duration.Seconds(), attrs)
}

func (t *Telemetry) Shutdown(ctx context.Context) error {
	if t.MeterProvider != nil {
		return t.MeterProvider.Shutdown(ctx)
	}
	return nil
}

