// Package swagui simplifies serving an instance of Swagger-UI. It can be added
// to a multiplexer, or served directly. If using a multiplexer, the path
// prefix option must match the relevant route.
package swagui

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"
)

var (
	urlKey = "url"
)

// Options holds optional Swagui data.
type Options struct {
	Version         int
	PathPrefix      string
	DefaultURLParam string
	NotFoundHandler http.Handler
}

// Swagui provides a Swagger-UI http.Handler and related data.
type Swagui struct {
	version         int
	prefix          string
	defaultURLParam string
	notFoundHandler http.Handler
	finder          assetFinder
	modtime         time.Time
}

// New returns a Swagui and defaults to the latest version of Swagger. If a
// path prefix is provided, it will be filtered for appropriate usage.
func New(opts *Options) (*Swagui, error) {
	if opts == nil {
		opts = &Options{}
	}

	s := &Swagui{
		version:         opts.Version,
		prefix:          filterPrefix(opts.PathPrefix),
		defaultURLParam: opts.DefaultURLParam,
		notFoundHandler: opts.NotFoundHandler,
		modtime:         time.Now(),
	}

	if s.notFoundHandler == nil {
		s.notFoundHandler = http.NotFoundHandler()
	}

	switch s.version {
	case 1:
		s.finder = &data1{}
	default:
		s.finder = &data2{}
	}

	return s, nil
}

// PathPrefix returns the current path prefix which has been filtered for
// appropriate usage.
func (s *Swagui) PathPrefix() string {
	return s.prefix
}

func (s *Swagui) urlParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.defaultURLParam == "" || r.URL.Path != s.prefix {
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

		u := "./?" + urlKey + "=" + s.defaultURLParam + rq

		http.Redirect(w, r, u, 301)
	})
}

func (s *Swagui) handler(w http.ResponseWriter, r *http.Request) {
	// redirect "root" requests missing trailing slash
	if len(r.URL.Path) > 1 && r.URL.Path == s.prefix[:len(s.prefix)-1] {
		http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
		return
	}

	name := strings.TrimPrefix(r.URL.Path, s.prefix)

	// rename index requests
	if name == "" {
		name = "index.html"
	}

	b, err := filteredAsset(s.finder, name)
	if err != nil {
		s.notFoundHandler.ServeHTTP(w, r)
		return
	}

	c := bytes.NewReader(b)
	http.ServeContent(w, r, path.Base(name), s.modtime, c)
}

// Handler returns an http.Handler which serves Swagger-UI.
func (s *Swagui) Handler() http.Handler {
	return s.urlParam(http.HandlerFunc(s.handler))
}

func filterPrefix(p string) string {
	p = path.Clean(p)

	// ensure path.Clean dot return is avoided
	if p == "." {
		p = ""
	}

	// ensure prefix begins with slash
	if p == "" || p[0] != '/' {
		p = "/" + p
	}

	// ensure prefix ends with slash
	if len(p) > 1 && p[len(p)-1] != '/' {
		p += "/"
	}

	return p
}

func filteredAsset(finder assetFinder, name string) ([]byte, error) {
	if name == "swagger-ui.js" {
		name = "swagger-ui.min.js"
	}

	return finder.Asset(name)
}
