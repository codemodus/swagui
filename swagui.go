// Package swagui simplifies serving an instance of Swagger-UI.
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

// Swagui provides a Swagger-UI http.Handler and manages related data.
type Swagui struct {
	fs        *uiFiles
	nfHandler http.Handler
	modtime   time.Time
	indexFile string
}

// New returns a Swagui and defaults to the latest version of Swagger.
func New(opts *Options) (*Swagui, error) {
	if opts == nil {
		opts = &Options{}
	}

	s := Swagui{
		fs:        newUIFiles(opts.Version),
		nfHandler: opts.NotFoundHandler,
		modtime:   time.Now(),
		indexFile: "index.html",
	}

	if s.nfHandler == nil {
		s.nfHandler = http.NotFoundHandler()
	}

	return &s, nil
}

// Handler returns an http.Handler which serves Swagger-UI.
func (s *Swagui) Handler(defaultDefinition string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p != "" && p[0] == '/' {
			p = p[1:]
		}

		var bs []byte
		var err error

		switch p {
		case "", s.indexFile:
			bs, err = s.fs.defFilteredFile(s.indexFile, defaultDefinition)
		default:
			bs, err = s.fs.file(p)
		}
		if err != nil {
			s.nfHandler.ServeHTTP(w, r)
			return
		}

		http.ServeContent(w, r, p, s.modtime, bytes.NewReader(bs))
	})
}
