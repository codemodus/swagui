package suihttp

import (
	"bytes"
	"net/http"
	"time"

	"github.com/codemodus/swagui/internal/assets"
)

// Handler returns an http.Handler which serves resources using special handling
// for Swagger-UI.
func Handler(as *assets.Assets, notFound http.Handler, index, origDef, newDef string) http.Handler {
	modtime := time.Now()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if p != "" && p[0] == '/' {
			p = p[1:]
		}

		var bs []byte
		var err error

		switch p {
		case "", index:
			bs, err = as.ModifiedAsset(index, origDef, newDef)
		default:
			bs, err = as.Asset(p)
		}
		if err != nil {
			notFound.ServeHTTP(w, r)
			return
		}

		http.ServeContent(w, r, p, modtime, bytes.NewReader(bs))
	})
}
