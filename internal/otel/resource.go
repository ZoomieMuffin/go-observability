package otel

import (
	"context"
	"os"

	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func NewResource(defaultServiceName string) (*sdkresource.Resource, error) {
	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName == "" {
		serviceName = defaultServiceName
	}

	return sdkresource.New(
		context.TODO(),
		sdkresource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
}
