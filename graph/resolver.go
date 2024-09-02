//go:generate go run github.com/99designs/gqlgen generate
package graph

import example "github.com/deitrix/gqlgen-dataloadgen-example"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store example.Store
}

func NewResolver(store example.Store) *Resolver {
	return &Resolver{
		store: store,
	}
}
