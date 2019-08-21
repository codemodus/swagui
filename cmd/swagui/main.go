package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codemodus/swagui"
	"github.com/codemodus/swagui/suidata1"
	"github.com/codemodus/swagui/suidata2"
	"github.com/codemodus/swagui/suidata3"
)

func main() {
	var (
		port    = ":2288"
		def     = ""
		scrape  = ""
		version = 0
		scrpath = "/swagger.json"
	)

	flag.StringVar(&port, "port", port, "http port")
	flag.StringVar(&def, "def", def, "remote definition")
	flag.StringVar(&scrape, "scrape", scrape, "definition to scrape and serve")
	flag.IntVar(&version, "v", version, "swagger version")
	flag.Parse()

	defType := "default"
	defSrc := def
	if scrape != "" {
		if def != "" {
			fmt.Printf("overriding def flag\n")
		}
		def = fmt.Sprintf("http://localhost%s%s", port, scrpath)
		defType = "scraped"
		defSrc = scrape
	}

	var p swagui.Provider
	switch version {
	case 1:
		p = suidata1.New()
	case 2:
		p = suidata2.New()
	default:
		version = 3
		p = suidata3.New()
	}
	ui, err := swagui.New(http.NotFoundHandler(), p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	msgfmt := "serving swaggerui v%d on %s, with %s def %q"
	fmt.Printf(msgfmt, version, port, defType, defSrc)

	var h http.Handler = ui.Handler(def)
	wrap, err := scrapeHandlerFunc(scrpath, scrape)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	h = wrap(h)

	if err = http.ListenAndServe(port, h); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func scrapeHandlerFunc(path, resource string) (func(http.Handler) http.Handler, error) {
	if resource == "" {
		return passthroughHandler, nil
	}

	data, err := getData(resource)
	if err != nil {
		return nil, err
	}

	return jsonHandlerFunc(path, data), nil
}

func passthroughHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func jsonHandlerFunc(path string, data []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != path {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		})
	}
}

func getData(resource string) ([]byte, error) {
	cl := http.Client{Timeout: time.Second * 60}

	res, err := cl.Get(resource)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bb := &bytes.Buffer{}
	if _, err = bb.ReadFrom(res.Body); err != nil {
		return nil, err
	}

	return bb.Bytes(), nil
}
