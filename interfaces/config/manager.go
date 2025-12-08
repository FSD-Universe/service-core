// Package config
package config

type Verifiable interface {
	Verify() (bool, error)
}

type DefaultSetter interface {
	InitDefaults()
}

type Item interface {
	DefaultSetter
	Verifiable
}

type ManagerInterface[T Item] interface {
	Init() error
	GetConfig() T
	SaveConfig() error
}
