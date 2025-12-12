// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build event_bus

// Package bus
package bus

import (
	"context"
)

type Event[T comparable] struct {
	Type T
	Data interface{}
}

type Subscriber[T comparable] func(message *Event[T]) error

type EventBusInterface[T comparable] interface {
	Start(ctx context.Context)
	Stop()
	Publish(event *Event[T])
	Subscribe(eventType T, handler Subscriber[T])
	Shutdown(ctx context.Context) error
}
