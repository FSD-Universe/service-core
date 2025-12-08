//go:build database

// Package entity
package entity

type Base interface {
	comparable
	GetId() uint
	SetId(id uint)
}

type Comparable[T Base] interface {
	Equal(other T) bool
	Diff(other T) map[string]interface{}
}
