// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build telemetry

// Package telemetry
package telemetry

import (
	"context"
	"errors"

	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	t "go.opentelemetry.io/otel/trace"
)

type SDK struct {
	logger logger.Interface
	config *config.TelemetryConfig
	tracer t.Tracer
}

func NewSDK(lg logger.Interface, c *config.TelemetryConfig) *SDK {
	return &SDK{
		logger: logger.NewLoggerAdapter(lg, "telemetry"),
		config: c,
		tracer: otel.Tracer(c.Name),
	}
}

func (sdk *SDK) SetupOTelSDK(ctx context.Context) (shutdown cleaner.ShutdownCallback, err error) {
	var shutdownFuncs []cleaner.ShutdownCallback

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	tracerProvider, err := sdk.NewTraceProvider(ctx)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)
	return
}

func (sdk *SDK) NewTraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(sdk.config.Endpoint),
		otlptracegrpc.WithInsecure(),
	}
	exporter, err := otlptracegrpc.New(ctx, opts...)
	if err != nil {
		sdk.logger.Fatalf("failed to create new trace provider: %v", err)
	}

	r, err := resource.New(ctx, []resource.Option{
		resource.WithAttributes(
			attribute.KeyValue{Key: "token", Value: attribute.StringValue(sdk.config.Token)},
			attribute.KeyValue{Key: "service.name", Value: attribute.StringValue(sdk.config.ServiceName)},
			attribute.KeyValue{Key: "host.name", Value: attribute.StringValue(sdk.config.HostName)},
		),
	}...)
	if err != nil {
		sdk.logger.Fatalf("failed to create new resource: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func InitSDK(lg logger.Interface, cl cleaner.Interface, c *config.TelemetryConfig) error {
	sdk := NewSDK(lg, c)
	shutdown, err := sdk.SetupOTelSDK(context.Background())
	if err != nil {
		lg.Fatalf("fail to initialize telemetry: %v", err)
		return err
	}
	cl.Add("Telemetry", shutdown)
	return nil
}
