package suidata3

import (
	"net/http"

	"github.com/codemodus/swagui/internal/assets"
	"github.com/codemodus/swagui/internal/suihttp"
)

// SUIData3 represents Swagger-UI v3 resources.
type SUIData3 struct {
	as    *assets.Assets
	index string
	def   string
}

// New sets up a new SUIData3.
func New() *SUIData3 {
	aliases := map[string]string{}

	return &SUIData3{
		as:    assets.New(Asset, aliases),
		index: "index.html",
		def:   "https://petstore.swagger.io/v2/swagger.json",
	}
}

// Handler returns an http.Handler which serves Swagger-UI.
func (d *SUIData3) Handler(notFound http.Handler, defaultDef string) http.Handler {
	return suihttp.Handler(d.as, notFound, d.index, d.def, defaultDef)
}
