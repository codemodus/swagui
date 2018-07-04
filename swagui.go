// Package swagui simplifies serving an instance of Swagger-UI. It can be added
// to a multiplexer, or served directly. If using a multiplexer, the path
// prefix option must match the relevant route.
package swagui

import (
	"bytes"
	"net/http"
	"time"
)

// Options holds optional Swagui data.
type Options struct {
	Version         Version
	NotFoundHandler http.Handler
}

// Swagui provides a Swagger-UI http.Handler and related data.
type Swagui struct {
	files           *uiFiles
	notFoundHandler http.Handler
	modtime         time.Time
}

// New returns a Swagui and defaults to the latest version of Swagger. If a
// path prefix is provided, it will be filtered for appropriate usage.
func New(opts *Options) (*Swagui, error) {
	if opts == nil {
		opts = &Options{}
	}

	s := &Swagui{
		files:           newUIFiles(opts.Version),
		notFoundHandler: opts.NotFoundHandler,
		modtime:         time.Now(),
	}

	if s.notFoundHandler == nil {
		s.notFoundHandler = http.NotFoundHandler()
	}

	return s, nil
}

// Handler returns an http.Handler which serves Swagger-UI.
func (s *Swagui) Handler(defaultDefinition string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p != "" && p[0] == '/' {
			p = p[1:]
		}

		var b []byte
		var err error

		switch p {
		case "", indexFile:
			b, err = s.files.accessIndex(defaultDefinition)
		default:
			b, err = s.files.access(p)
		}
		if err != nil {
			s.notFoundHandler.ServeHTTP(w, r)
			return
		}

		c := bytes.NewReader(b)
		http.ServeContent(w, r, p, s.modtime, c)
	})
}
