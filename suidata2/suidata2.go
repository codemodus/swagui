package suidata2

import (
	"net/http"

	"github.com/codemodus/swagui/internal/assets"
	"github.com/codemodus/swagui/internal/suihttp"
)

// SUIData2 represents Swagger-UI v2 resources.
type SUIData2 struct {
	as    *assets.Assets
	index string
	def   string
}

// New sets up a new SUIData2.
func New() *SUIData2 {
	aliases := map[string]string{
		"swagger-ui.js": "swagger-ui.min.js",
	}

	as := assets.New(Asset, aliases)

	return &SUIData2{
		as:    as,
		index: "index.html",
		def:   "http://petstore.swagger.io/v2/swagger.json",
	}
}

// Handler returns an http.Handler which serves Swagger-UI.
func (d *SUIData2) Handler(notFound http.Handler, defaultDef string) http.Handler {
	return suihttp.Handler(d.as, notFound, d.index, d.def, defaultDef)
}
