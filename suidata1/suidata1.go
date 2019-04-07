package suidata1

import (
	"net/http"

	"github.com/codemodus/swagui/internal/assets"
	"github.com/codemodus/swagui/internal/suihttp"
)

// SUIData1 represents Swagger-UI v1 resources.
type SUIData1 struct {
	as    *assets.Assets
	index string
	def   string
}

// New sets up a new SUIData1.
func New() *SUIData1 {
	aliases := map[string]string{
		"swagger-ui.js": "swagger-ui.min.js",
	}

	return &SUIData1{
		as:    assets.New(Asset, aliases),
		index: "index.html",
		def:   "http://petstore.swagger.wordnik.com/api/api-docs.json",
	}
}

// Handler returns an http.Handler which serves Swagger-UI.
func (d *SUIData1) Handler(notFound http.Handler, defaultDef string) http.Handler {
	return suihttp.Handler(d.as, notFound, d.index, d.def, defaultDef)
}
