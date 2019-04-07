// Package swagui simplifies serving an instance of Swagger-UI.
package swagui

import (
	"fmt"
	"net/http"
)

// Provider sets up an HTTP handler for Swagger-UI resources. The provided types
// that implement this interface are located in separate packages based on their
// Swagger version so that imported dependencies are controlled by the caller.
type Provider interface {
	Handler(notFound http.Handler, defaultDef string) http.Handler
}

// Swagui wraps a Provider to simplify usage.
type Swagui struct {
	nf http.Handler
	p  Provider
}

// New returns a Swagui, or an error if no provider is set.
func New(notFound http.Handler, p Provider) (*Swagui, error) {
	efmt := "new swagui: %s"

	if notFound == nil {
		notFound = http.NotFoundHandler()
	}
	if p == nil {
		return nil, fmt.Errorf(efmt, "provider must be set")
	}

	s := Swagui{
		nf: notFound,
		p:  p,
	}

	return &s, nil
}

// Handler returns an http.Handler for Swagger-UI resources.
func (s *Swagui) Handler(defaultDefinition string) http.Handler {
	return s.p.Handler(s.nf, defaultDefinition)
}
