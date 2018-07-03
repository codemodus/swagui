// Package swagui simplifies serving an instance of Swagger-UI. It can be added
// to a multiplexer, or served directly. If using a multiplexer, the path
// prefix option must match the relevant route.
package swagui

import (
	"bytes"
	"net/http"
	"time"
)

var (
	urlKey = "url"
	index  = []byte("index.html")
)

// Version wraps integer to ease the defining of the swagger ui version.
type Version int

// V{N} constants list the available swagger ui versions.
const (
	V1 Version = iota
	V2
	V3
)

// Options holds optional Swagui data.
type Options struct {
	Version         Version
	NotFoundHandler http.Handler
}

// Swagui provides a Swagger-UI http.Handler and related data.
type Swagui struct {
	notFoundHandler http.Handler
	accessor        accessor
	modtime         time.Time
}

// New returns a Swagui and defaults to the latest version of Swagger. If a
// path prefix is provided, it will be filtered for appropriate usage.
func New(opts *Options) (*Swagui, error) {
	if opts == nil {
		opts = &Options{}
	}

	s := &Swagui{
		notFoundHandler: opts.NotFoundHandler,
		accessor:        accessorByVersion(opts.Version),
		modtime:         time.Now(),
	}

	if s.notFoundHandler == nil {
		s.notFoundHandler = http.NotFoundHandler()
	}

	return s, nil
}

func setDefaultDefinition(def string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if def == "" || r.URL.Path != "" && r.URL.Path != "/" {
				next.ServeHTTP(w, r)
				return
			}

			p := r.URL.Query().Get(urlKey)
			if p != "" {
				next.ServeHTTP(w, r)
				return
			}

			rq := r.URL.RawQuery
			if rq != "" {
				rq = "&" + rq
			}

			u := "./?" + urlKey + "=" + def + rq

			http.Redirect(w, r, u, 301)
		})
	}
}

func (s *Swagui) handler(w http.ResponseWriter, r *http.Request) {
	p := []byte(r.URL.Path)
	if len(p) > 0 && p[0] == '/' {
		p = p[1:]
	}

	if len(p) == 0 {
		p = index
	}

	path := string(p)

	b, err := s.accessor.Access(path)
	if err != nil {
		s.notFoundHandler.ServeHTTP(w, r)
		return
	}

	c := bytes.NewReader(b)
	http.ServeContent(w, r, path, s.modtime, c)
}

// Handler returns an http.Handler which serves Swagger-UI.
func (s *Swagui) Handler(defaultDefinition string) http.Handler {
	return setDefaultDefinition(defaultDefinition)(http.HandlerFunc(s.handler))
}
